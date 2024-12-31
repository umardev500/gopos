package container

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/app/handler"
	"gitub.com/umardev500/gopos/app/repository"
	"gitub.com/umardev500/gopos/app/service"
	"gitub.com/umardev500/gopos/internal/app/contract"
	pkgContract "gitub.com/umardev500/gopos/pkg/contract"
	"gitub.com/umardev500/gopos/pkg/database"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type authContainer struct {
	authHandler contract.AuthHandler
}

func NewAuthContainer(db *database.GormInstance, v validator.Validator) pkgContract.Container {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, v)
	authHandler := handler.NewAuthHandler(authService)

	return &authContainer{
		authHandler: authHandler,
	}
}

func (a *authContainer) HandleApi(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", a.authHandler.Login)
}

func (a *authContainer) HandleWeb(router fiber.Router) {}
