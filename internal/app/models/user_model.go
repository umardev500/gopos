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

type CreateUserRequest struct {
	ID           string   `json:"-"`
	Username     string   `json:"username" validate:"required,min=6"` // unique
	Email        string   `json:"email" validate:"required,email"`    // unique
	Password     string   `gorm:"column:password_hash" json:"password" validate:"required,min=6"`
	PasswordConf string   `gorm:"-" json:"password_conf" validate:"required,min=6,eqfield=Password"`
	Roles        []string `gorm:"-" json:"role_ids" validate:"required,min=1"`
}

func (CreateUserRequest) TableName() string {
	return "users"
}

type UpdateUserRequest struct {
	ID string `json:"-"`
}

type FindUsersParams struct {
	TenantID   *string
	Pagination pkgModel.PaginationParams
}
