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
	// load .env from file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// connect to db
	user, passwd, dbname := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	credentials := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, passwd, dbname)
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// seed db
	if err := postgres.Seeder(db); err != nil {
		panic(err)
	}

	// create service
	ts := &postgres.TaskService{DB: db}

	// embed service to handler
	// needs to implement service interface
	var h http.Handler
	h.TaskService = ts

	var r http.Router
	r.InitRoutes(h)

	r.Run(h)
}
