package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TaskService interface {
	Get(id int) (*Task, error)
	GetAll() ([]*Task, error)
	Create(t *Task) error
	Delete(id int) error
	Update(t Task) error
}
