package model

import "time"

type Task struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
