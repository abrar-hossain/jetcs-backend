package models

import "time"

type Author struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password,omitempty" binding:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
}
