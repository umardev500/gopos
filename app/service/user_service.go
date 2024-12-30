package service

import (
	"context"

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
	var resp = pkgUtil.Response{}

	paginatedResult, err := s.repo.GetAllUsers(ctx, params)
	if err != nil {
		return &resp
	}

	resp.Success = true
	resp.Message = "Find all users"
	resp.Data = paginatedResult.Data

	return &resp
}
