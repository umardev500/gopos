package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
)

type authHandler struct {
	authService contract.AuthService
}

func NewAuthHandler(authService contract.AuthService) contract.AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var payload models.LoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	ctx, cancel := pkgUtil.BaseContext()
	defer cancel()

	res := h.authService.Login(ctx, &payload)

	return c.Status(res.StatusCode).JSON(res)
}
