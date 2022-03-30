package postgres

import (
	"database/sql"
	app "go-api-test-2/models"
)

type TaskService struct {
	DB *sql.DB
}

func (s *TaskService) Tasks() ([]*app.Task, error) {
	var ts []*app.Task

	rows, err := s.DB.Query("SELECT * FROM tasks;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &app.Task{}
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.IsDone)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
