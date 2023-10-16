package database

import (
	"database/sql"
	"fmt"
)

func StartConnection() (*sql.DB, error) {
	schemaURL := fmt.Sprintf("%s@tcp(%s)/%s?parseTime=true", "root", "localhost:3306", "task_management")

	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//create migrate file
//migrate create -ext sql -dir db/migrations/users create_table_users
//migrate create -ext sql -dir db/migrations/tasks create_table_tasks

//migrate up:
//migrate -database "mysql://root@tcp(localhost:3306)/task_management" -path db/migrations/users up
//migrate -database "mysql://root@tcp(localhost:3306)/task_management" -path db/migrations/tasks up