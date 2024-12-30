package container

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/app/handler"
	"gitub.com/umardev500/gopos/app/repository"
	"gitub.com/umardev500/gopos/app/service"
	"gitub.com/umardev500/gopos/internal/app/contract"
	pgkContract "gitub.com/umardev500/gopos/pkg/contract"
	"gitub.com/umardev500/gopos/pkg/database"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type userContainer struct {
	userHandler contract.UserHandler
}

func NewUserContainer(db *database.GormInstance, v validator.Validator) pgkContract.Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &userContainer{
		userHandler: userHandler,
	}
}

func (u *userContainer) HandleApi(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", u.userHandler.GetAllUsers)
}

func (u *userContainer) HandleWeb(router fiber.Router) {}
