package repository

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/database"
)

type authRepository struct {
	db *database.GormInstance
}

func NewAuthRepository(db *database.GormInstance) contract.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (a *authRepository) Login(ctx context.Context, username string) (*models.AuthUser, error) {
	conn := a.db.GetConn(ctx)
	var user models.AuthUser

	err := conn.
		Select("users.id", "users.username", "users.email", "users.password_hash", "ut.tenant_id", "ub.branch_id").
		Joins("LEFT JOIN user_tenants ut ON ut.user_id = users.id").
		Joins("LEFT JOIN user_branches ub ON ub.user_id = users.id").
		Where("users.username = ?", username).
		Or("users.email = ?", username).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
