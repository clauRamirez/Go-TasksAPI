package http

import (
	"fmt"
	"github.com/gorilla/mux"
	app "go-api-test-2/models"
	"net/http"
)

type Handler struct {
	TaskService app.TaskService
}

func (h *Handler) InitRoutes(r *mux.Router) {
	r.HandleFunc("/test", func(w http.ResponseWriter, r2 *http.Request) {
		fmt.Fprint(w, "Test")
	}).Methods("GET")

	r.HandleFunc("/tasks", h.tasks).Methods("GET")
}

func (h *Handler) Run(r *mux.Router) {
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
