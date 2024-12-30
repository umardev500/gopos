package service

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
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

	paginatedResult, err := s.repo.GetAllUsers(ctx, params)
	if err != nil {
		return &pkgUtil.Response{}
	}

	page := params.Pagination.Page
	limit := params.Pagination.Limit
	data := paginatedResult.Data
	total := paginatedResult.Total
	totalPage := math.Ceil(float64(total) / float64(limit))

	return &pkgUtil.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Find all users",
		Data:       data,
		Pagination: &pkgUtil.PaginationMeta{
			Page:      int64(page),
			PageSize:  int64(limit),
			Total:     int64(total),
			TotalPage: int64(totalPage),
		},
	}
}
