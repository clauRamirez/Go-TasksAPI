package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TaskService interface {
	//Task(id int) (*Task, error)
	Tasks() ([]*Task, error)
	//CreateTask(u *Task) error
	//DeleteTask(id int) error
	//UpdateTask(u *Task) error
}
