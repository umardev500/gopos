package service

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/constant"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
)

type userService struct {
	repo contract.UserRepository
}

func NewUserService(repo contract.UserRepository) contract.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers(ctx context.Context, params *models.FindUsersParams) *pkgUtil.Response {
	claims := ctx.Value(constant.ClaimsContextKey).(jwt.MapClaims)

	// Assign user data
	if tid, ok := claims["tid"].(string); ok {
		params.TenantID = &tid
	}

	// Get all users
	paginatedResult, err := s.repo.GetAllUsers(ctx, params)
	if err != nil {
		return &pkgUtil.Response{}
	}

	// Assign pagination meta
	page := params.Pagination.Page
	limit := params.Pagination.Limit
	data := paginatedResult.Data
	total := paginatedResult.Total
	totalPage := math.Ceil(float64(total) / float64(limit))
	pagination := &pkgUtil.PaginationMeta{
		Page:      int64(page),
		PageSize:  int64(limit),
		Total:     int64(total),
		TotalPage: int64(totalPage),
	}

	// Return response
	return &pkgUtil.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Find all users",
		Data:       data,
		Pagination: pagination,
	}
}
