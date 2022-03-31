package http

import (
	"go-api-test-2/models"
)

/*
	Handles user request/responses to/from API
*/

type Handler struct {
	TaskService models.TaskService
}
