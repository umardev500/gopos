package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/contract"
	pkgModel "gitub.com/umardev500/gopos/pkg/model"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
)

type userHandler struct {
	userService contract.UserService
}

func NewUserHandler(userService contract.UserService) contract.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) GetAllUsers(c *fiber.Ctx) error {
	var paginationParams = &pkgModel.PaginationParams{
		Page:  c.QueryInt("page"),
		Limit: c.QueryInt("limit"),
	}

	ctx, cancel := pkgUtil.BaseContext()
	defer cancel()

	res := u.userService.GetAllUsers(ctx, paginationParams)
	if res == nil {
		return c.JSON(res)
	}

	return c.JSON(res)
}
