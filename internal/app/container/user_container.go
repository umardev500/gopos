package container

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/app/handler"
	"gitub.com/umardev500/gopos/app/repository"
	"gitub.com/umardev500/gopos/app/service"
	"gitub.com/umardev500/gopos/internal/app/contract"
	pgkContract "gitub.com/umardev500/gopos/pkg/contract"
	"gitub.com/umardev500/gopos/pkg/database"
	"gitub.com/umardev500/gopos/pkg/middleware"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type userContainer struct {
	userHandler contract.UserHandler
}

func NewUserContainer(db *database.GormInstance, v validator.Validator) pgkContract.Container {
	userRepo := repository.NewUserRepository(db)
	userTenantRepo := repository.NewUserTenantRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	userService := service.NewUserService(userRepo, userTenantRepo, roleRepo, userRoleRepo, db, v)
	userHandler := handler.NewUserHandler(userService)

	return &userContainer{
		userHandler: userHandler,
	}
}

func (u *userContainer) HandleApi(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", middleware.AuthMiddleware(), u.userHandler.GetAllUsers)
	users.Post("/", middleware.AuthMiddleware(), u.userHandler.CreateUser)
}

func (u *userContainer) HandleWeb(router fiber.Router) {}
