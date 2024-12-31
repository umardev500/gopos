package service

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/constant"
	"gitub.com/umardev500/gopos/pkg/database"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type userService struct {
	repo     contract.UserRepository
	db       *database.GormInstance
	validate validator.Validator
}

func NewUserService(repo contract.UserRepository, db *database.GormInstance, v validator.Validator) contract.UserService {
	return &userService{
		repo:     repo,
		db:       db,
		validate: v,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.CreateUserRequest) *pkgUtil.Response {
	// Validate input payload
	res := s.validate.Struct(user).Response()
	if res != nil {
		return res
	}

	// Start transaction to insert to user and user_roles
	err := s.db.WithTransaction(ctx, func(ctx context.Context) error {
		// Create user
		err := s.repo.CreateUser(ctx, user)

		// TODO: Create user roles

		return err
	})
	if err != nil {
		return &pkgUtil.Response{}
	}

	return &pkgUtil.Response{
		Success: true,
		Message: "User created successfully",
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
