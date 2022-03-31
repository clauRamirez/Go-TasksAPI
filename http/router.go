package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*
	HTTP server
*/
type Router struct {
	mux.Router
}

func (r *Router) InitRoutes(h Handler) {
	sr := r.StrictSlash(true).PathPrefix("/api").Subrouter()

	sr.HandleFunc("/tasks", h.tasks).Methods("GET")
	sr.HandleFunc("/tasks", h.createTask).Methods("POST")
	sr.HandleFunc("/tasks/{id}", h.task).Methods("GET")
	sr.HandleFunc("/tasks/{id}", h.deleteTask).Methods("DELETE")
	sr.HandleFunc("/tasks/{id}", h.updateTask).Methods("PUT")
}

func (r *Router) Run(addr string) {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", addr), r); err != nil {
		panic(err)
	}
}
