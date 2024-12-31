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

func (u *userHandler) CreateUser(c *fiber.Ctx) error {
	var payload models.CreateUserRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	ctx := pkgUtil.NewContext().WithTimeout(5).WithClaims(c)
	defer ctx.Cancel()

	res := u.userService.CreateUser(ctx.Ctx, &payload)

	return c.Status(res.StatusCode).JSON(res)
}

func (u *userHandler) GetAllUsers(c *fiber.Ctx) error {
	var paginationParams = &pkgModel.PaginationParams{
		Page:  c.QueryInt("page"),
		Limit: c.QueryInt("limit"),
	}

	ctx := pkgUtil.NewContext().WithTimeout(5).WithClaims(c)
	defer ctx.Cancel()

	// Add the claims to the context
	ctx.Ctx = context.WithValue(ctx.Ctx, constant.ClaimsContextKey, c.Locals(constant.ClaimsContextKey))

	params := models.FindUsersParams{
		Pagination: *paginationParams.Parse(),
	}

	res := u.userService.GetAllUsers(ctx.Ctx, &params)
	if res == nil {
		return c.JSON(res)
	}

	return c.JSON(res)
}
