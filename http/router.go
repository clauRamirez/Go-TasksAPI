package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux.Router
}

func (r *Router) InitRoutes(h Handler) {
	sr := r.StrictSlash(true).PathPrefix("/api").Subrouter()

	sr.HandleFunc("/test", func(w http.ResponseWriter, r2 *http.Request) {
		fmt.Fprint(w, "Test")
	}).Methods("GET")

	sr.HandleFunc("/tasks", h.tasks).Methods("GET")
}

func (r *Router) Run(h Handler) {
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
