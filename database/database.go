package database

import (
	"database/sql"
	"fmt"
)

func StartConnection() (*sql.DB, error) {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", "root", "root", "localhost:3306", "task_management")

	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
