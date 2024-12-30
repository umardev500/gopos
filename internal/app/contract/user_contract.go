package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/models"
	pkgModel "gitub.com/umardev500/gopos/pkg/model"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
)

type UserHandler interface {
	// Get all users
	GetAllUsers(c *fiber.Ctx) error
}

type UserService interface {
	// Get all users
	GetAllUsers(ctx context.Context, paginationParams *pkgModel.PaginationParams) *pkgUtil.Response
}

type UserRepository interface {
	// Create a new user
	CreateUser(ctx context.Context, user models.CreateUserRequest) error

	// Delete user by id
	DeleteUserById(ctx context.Context, id string) error

	// Delete multiple users with given ids
	DeleteUsers(ctx context.Context, ids []string) error

	// Get all users
	GetAllUsers(ctx context.Context, paginationParams *pkgModel.PaginationParams) (*pkgModel.PaginatedResult, error)

	// Get user by id
	GetUserById(ctx context.Context, id string) (*models.User, error)

	// Update user by id
	UpdateUserById(ctx context.Context, user models.UpdateUserRequest) error
}
