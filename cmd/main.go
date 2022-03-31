package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-api-test-2/http"
	"go-api-test-2/postgres"
	"os"
)

func main() {
	fmt.Println("Loading env variables...")
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	var PORT string
	if val, ok := os.LookupEnv("PORT"); ok {
		PORT = val
	} else {
		PORT = "8080"
	}

	fmt.Println("Connecting to database...")
	user, passwd, dbname := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	credentials := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, passwd, dbname)
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Seeding database...")
	if err := postgres.Seeder(db); err != nil {
		panic(err)
	}

	fmt.Println("Creating services and embedding handlers...")
	// create service
	ts := &postgres.TaskService{DB: db}

	// embed service to handler
	// needs to implement service interface
	var h http.Handler
	h.TaskService = ts

	fmt.Println("Initializing HTTP routes...")
	var r http.Router
	r.InitRoutes(h)

	fmt.Printf("HTTP server listening on port %s\n", PORT)
	r.Run(PORT)
}
