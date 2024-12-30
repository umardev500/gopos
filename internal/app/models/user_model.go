package models

import (
	"github.com/google/uuid"
	pkgModel "gitub.com/umardev500/gopos/pkg/model"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Version      int       `json:"version"`

	pkgModel.Time
}

type CreateUserRequest struct{}

type UpdateUserRequest struct {
	ID string `json:"-"`
}
