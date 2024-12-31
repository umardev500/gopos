package service

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/constant"
	"gitub.com/umardev500/gopos/pkg/database"
	pkgUtil "gitub.com/umardev500/gopos/pkg/util"
	"gitub.com/umardev500/gopos/pkg/validator"
)

type userService struct {
	repo           contract.UserRepository
	userTenantRepo contract.UserTenantRepository
	roleRepo       contract.RoleRepository
	userRoleRepo   contract.UserRoleRepository
	db             *database.GormInstance
	validate       validator.Validator
}

func NewUserService(
	repo contract.UserRepository,
	userTenantRepo contract.UserTenantRepository,
	roleRepo contract.RoleRepository,
	userRoleRepo contract.UserRoleRepository,
	db *database.GormInstance,
	v validator.Validator,
) contract.UserService {
	return &userService{
		repo:           repo,
		userTenantRepo: userTenantRepo,
		roleRepo:       roleRepo,
		userRoleRepo:   userRoleRepo,
		db:             db,
		validate:       v,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.CreateUserRequest) *pkgUtil.Response {
	// Validate input payload
	res := s.validate.Struct(user).Response()
	if res != nil {
		return res
	}

	userID := uuid.New().String()
	user.ID = userID
	var tenantID *string

	// Assign tenant id from context
	claims := ctx.Value(constant.ClaimsContextKey).(jwt.MapClaims)
	if tid, ok := claims["tid"].(string); ok {
		tenantID = &tid
	}

	// count given roles
	roleCount, err := s.roleRepo.CountRolesByTenantID(ctx, user.Roles, tenantID)
	if err != nil {
		return pkgUtil.DBErrorResponse(err)
	}

	// compare it does the roles give is valid with related tenant id
	// if tenant id is not nil
	// if nil that indicate platform query
	if roleCount != int64(len(user.Roles)) {
		return pkgUtil.BadRequestResponse(constant.ErrRoleCountNotValid)
	}

	// Start transaction to insert to user and user_roles
	err = s.db.WithTransaction(ctx, func(ctx context.Context) error {
		var err error

		// Create user
		err = s.repo.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		// Mapping user roles
		var roles []*models.UserRoleParam
		for _, role := range user.Roles {
			roles = append(roles, &models.UserRoleParam{
				UserID: userID,
				RoleID: role,
			})
		}

		// Assign roles to the user
		err = s.userRoleRepo.AssignUserRoles(ctx, roles)
		if err != nil {
			return err
		}

		// Assign user to tenant if tenant id is provided
		// if provided that indicate that user created by tenant admin
		// else that indicate that user created by internal platform
		if tenantID != nil {
			err = s.userTenantRepo.AssignUserToTenant(ctx, &models.UserTenant{
				UserID:   userID,
				TenantID: *tenantID,
			})
		}

		return err
	})
	if err != nil {
		return pkgUtil.DBErrorResponse(err)
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
		return pkgUtil.DBErrorResponse(err)
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
