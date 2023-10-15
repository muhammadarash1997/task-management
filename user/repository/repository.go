package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhammadarash1997/task-management/user/model"

	"fmt"
)

var (
	insertUser = `
		INSERT INTO users(name, email, password) VALUES (?,?,?)
	`
	selectUsers = `
		SELECT * FROM users
	`
	selectUserById = `
		SELECT * FROM users WHERE id = ?
	`
	selectUserByEmail = `
		SELECT * FROM users WHERE email = ?
	`
	updateUser = `
		UPDATE users
		SET name = ?, email = ?
		WHERE id = ?
	`
	deleteUser = `
		DELETE FROM users
		WHERE id = ?
	`
)

type Repository interface {
	CreateUser(model.User) error
	GetAllUser() ([]model.User, error)
	GetUserById(uint) (model.User, error)
	GetUserByEmail(string) (model.User, error)
	UpdateUser(model.User) error
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

func (r *repository) CreateUser(user model.User) error {
	valueArgs := []interface{}{}
	valueArgs = append(valueArgs, user.Name)
	valueArgs = append(valueArgs, user.Email)
	valueArgs = append(valueArgs, user.Password)

	_, err := r.db.Exec(insertUser, valueArgs...)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}

func (r *repository) GetAllUser() ([]model.User, error) {
	users := []model.User{}
	rows, err := r.db.Query(selectUsers)
	if err != nil {
		return []model.User{}, fmt.Errorf("r.db.Query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return []model.User{}, fmt.Errorf("rows.Scan: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(selectUserByEmail, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("r.db.QueryRow: %v", err)
	}

	return user, nil
}

func (r *repository) GetUserById(id uint) (model.User, error) {
	user := model.User{}
	err := r.db.QueryRow(selectUserById, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("r.db.QueryRow: %v", err)
	}

	return user, nil
}

func (r *repository) UpdateUser(user model.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(updateUser, user.Name, user.Email, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}

func (r *repository) DeleteById(id uint) error {
	_, err := r.db.Exec(deleteUser, id)
	if err != nil {
		return fmt.Errorf("r.db.Exec: %v", err)
	}

	return nil
}
