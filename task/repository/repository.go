package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhammadarash1997/task-management/task/model"

	"fmt"
)

var (
	insertTask = `
		INSERT INTO tasks(user_id, title, description, status) VALUES (?,?,?,?)
	`
	selectTasks = `
		SELECT * FROM tasks
	`
	selectTaskById = `
		SELECT * FROM tasks WHERE id = ?
	`
	updateTask = `
		UPDATE tasks
		SET user_id = ?, title = ?, description = ?, status = ?
		WHERE id = ?
	`
	deleteTask = `
		DELETE FROM tasks
		WHERE id = ?
	`
)

type Repository interface {
	CreateTask(model.Task) error
	GetAllTask() ([]model.Task, error)
	GetTaskById(uint) (model.Task, error)
	UpdateTask(model.Task) error
	DeleteById(uint) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTask(task model.Task) error {
	valueArgs := []interface{}{}
	valueArgs = append(valueArgs, task.UserID)
	valueArgs = append(valueArgs, task.Title)
	valueArgs = append(valueArgs, task.Description)
	valueArgs = append(valueArgs, task.Status)

	_, err := r.db.Exec(insertTask, valueArgs...)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}

func (r *repository) GetAllTask() ([]model.Task, error) {
	tasks := []model.Task{}
	rows, err := r.db.Query(selectTasks)
	if err != nil {
		return []model.Task{}, fmt.Errorf("r.db.Query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return []model.Task{}, fmt.Errorf("rows.Scan: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *repository) GetTaskById(id uint) (model.Task, error) {
	task := model.Task{}
	err := r.db.QueryRow(selectTaskById, id).Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		return model.Task{}, fmt.Errorf("r.db.QueryRow: %v", err)
	}

	return task, nil
}

func (r *repository) UpdateTask(task model.Task) error {
	task.UpdatedAt = time.Now()
	_, err := r.db.Exec(updateTask, task.UserID, task.Title, task.Description, task.Status, task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}

func (r *repository) DeleteById(id uint) error {
	_, err := r.db.Exec(deleteTask, id)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}
