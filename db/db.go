package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDb(config EnvDbConvig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.GetUsernameDb(), config.GetPasswordDb(), config.GetHostDb(), config.GetPortDb(), config.GetNameDb())

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestDbConnection(db *sql.DB) string {
	err := db.Ping()
	if err != nil {
		return err.Error()
	}
	return "Pong"
}
