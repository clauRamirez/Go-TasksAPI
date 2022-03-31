package http

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api-test-2/models"
	"net/http"
	"strconv"
)

func (h *Handler) tasks(w http.ResponseWriter, r *http.Request) {
	ts, err := h.TaskService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	respondWithJSON(w, http.StatusOK, ts)
}

func (h *Handler) task(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	t, err := h.TaskService.Get(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "item not found")
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, t)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	if err = h.TaskService.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	r.Body.Close()

	if err := h.TaskService.Create(&t); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, t)
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid item id")
		return
	}

	var t models.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid item payload")
		return
	}
	r.Body.Close()

	t.Id = id

	if err := h.TaskService.Update(t); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, t)
}
