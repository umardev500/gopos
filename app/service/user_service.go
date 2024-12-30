package service

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	pkgModel "gitub.com/umardev500/gopos/pkg/model"
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

func (s *userService) GetAllUsers(ctx context.Context, paginationParams *pkgModel.PaginationParams) *pkgUtil.Response {
	var resp = pkgUtil.Response{}

	paginatedResult, err := s.repo.GetAllUsers(ctx, paginationParams.Parse())
	if err != nil {
		return &resp
	}

	resp.Success = true
	resp.Message = "Find all users"
	resp.Data = paginatedResult.Data

	return &resp
}
