package http

import (
	"go-api-test-2/models"
)

/*
	User request/responses to/from API

*/
type Handler struct {
	TaskService models.TaskService
}
