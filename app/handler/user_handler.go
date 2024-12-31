package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/constant"
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

	// Add the claims to the context
	ctx = context.WithValue(ctx, constant.ClaimsContextKey, c.Locals(constant.ClaimsContextKey))

	params := models.FindUsersParams{
		Pagination: *paginationParams.Parse(),
	}

	res := u.userService.GetAllUsers(ctx, &params)
	if res == nil {
		return c.JSON(res)
	}

	return c.JSON(res)
}
