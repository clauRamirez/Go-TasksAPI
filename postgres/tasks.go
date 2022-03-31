package postgres

import (
	"database/sql"
	"go-api-test-2/models"
)

/*
	Database implementation
*/
type TaskService struct {
	*sql.DB
}

func (s *TaskService) GetAll() ([]*models.Task, error) {
	var ts []*models.Task

	rows, err := s.Query("SELECT * FROM tasks;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &models.Task{}
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.IsDone)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func (s *TaskService) Get(id int) (*models.Task, error) {
	t := &models.Task{}

	err := s.QueryRow("SELECT * FROM tasks WHERE id=$1",
		id).Scan(&t.Id, &t.Name, &t.Description, &t.IsDone)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *TaskService) Delete(id int) error {
	_, err := s.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}

func (s *TaskService) Create(t *models.Task) error {
	err := s.QueryRow("INSERT INTO tasks(name, description, is_done) VALUES($1, $2, $3) RETURNING id",
		t.Name, t.Description, t.IsDone).Scan(&t.Id)
	return err
}

func (s *TaskService) Update(t models.Task) error {
	_, err := s.Exec("UPDATE tasks SET name=$1, description=$2, is_done=$3 WHERE id=$4",
		t.Name, t.Description, t.IsDone, t.Id)
	return err
}
