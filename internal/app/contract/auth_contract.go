package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/models"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type AuthService interface {
	Login(ctx context.Context, payload *models.LoginRequest) *pkgUtil.Response
}

type AuthRepository interface {
	Login(ctx context.Context, username string) (*models.AuthUser, error)
}
