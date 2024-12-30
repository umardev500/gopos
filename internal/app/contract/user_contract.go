package contract

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/models"
)

type UserHandler interface{}

type UserService interface{}

type UserRepository interface {
	// Create a new user
	CreateUser(ctx context.Context, user models.CreateUserRequest) error
}
