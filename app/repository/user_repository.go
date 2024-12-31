package repository

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/database"

	pkgModel "gitub.com/umardev500/gopos/pkg/model"
)

type userRepository struct {
	db *database.GormInstance
}

func NewUserRepository(db *database.GormInstance) contract.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.CreateUserRequest) error {
	conn := r.db.GetConn(ctx)
	return conn.Create(&user).Error
}

func (r *userRepository) DeleteUserById(ctx context.Context, id string) error {
	conn := r.db.GetConn(ctx)
	return conn.Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepository) DeleteUsers(ctx context.Context, ids []string) error {
	conn := r.db.GetConn(ctx)
	return conn.Delete(&models.User{}, "id IN (?)", ids).Error
}

func (r *userRepository) GetAllUsers(ctx context.Context, params *models.FindUsersParams) (*pkgModel.PaginatedResult, error) {
	conn := r.db.GetConn(ctx)
	query := conn.Model(&models.User{})
	var totalCount int64

	// Base query is join with user tenant to get user info
	query.Joins("LEFT JOIN user_tenants ut ON ut.user_id = users.id")

	// If tenant id is provided, filter by tenant
	// else filter by tenant id being null indicating internal platform user doing the request
	if params.TenantID != nil {
		query.Where("ut.tenant_id = ?", params.TenantID)
	} else {
		query.Where("ut.tenant_id IS NULL")
	}

	// Count total
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	pagination := params.Pagination
	pageSize := pagination.Limit
	offset := (pagination.Page - 1) * pageSize

	var users []*models.User
	if err := query.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return &pkgModel.PaginatedResult{
		Data:  users,
		Total: int(totalCount),
	}, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	conn := r.db.GetConn(ctx)

	var user models.User
	if err := conn.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*models.User, error) {
	conn := r.db.GetConn(ctx)
	var user models.User

	if err := conn.First(&user, "username = ? OR email = ?", usernameOrEmail, usernameOrEmail).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUserById(ctx context.Context, user models.UpdateUserRequest) error {
	panic("implement me")
}
