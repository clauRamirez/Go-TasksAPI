package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api-test-2/models"
	"net/http"
	"strconv"
)

// TO-DO: use respondWithError instead of panicking on error

func (h *Handler) tasks(w http.ResponseWriter, r *http.Request) {
	ts, err := h.TaskService.GetAll()
	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusOK, ts)
}

func (h *Handler) task(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		panic(err)
	}

	t, err := h.TaskService.Get(id)
	respondWithJSON(w, http.StatusOK, t)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		panic(err)
	}

	err = h.TaskService.Delete(id)
	if err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"response": "success"})
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}
	r.Body.Close()

	if err := h.TaskService.Create(&t); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, t)
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		panic(err)
	}

	var t models.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}
	r.Body.Close()

	t.Id = id

	if err := h.TaskService.Update(t); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, t)
}
