package http

import (
	app "go-api-test-2/models"
)

type Handler struct {
	TaskService app.TaskService
}
