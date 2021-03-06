package postgres

import (
	"database/sql"
	"io/ioutil"
)

func Seeder(db *sql.DB) error {
	query, err := ioutil.ReadFile("postgres/seeds.sql")
	_, err = db.Exec(string(query))
	return err
}
