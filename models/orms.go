package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name" json:"name"`
  Email string   `db:"email" json:"email"`
}

type Task struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	Priority    int       `db:"priority" json:"priority"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
}
