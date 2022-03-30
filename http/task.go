package http

import (
	"net/http"
)

func (h *Handler) tasks(w http.ResponseWriter, r *http.Request) {
	ts, err := h.TaskService.Tasks()
	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusOK, ts)
}
