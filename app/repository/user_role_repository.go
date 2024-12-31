package repository

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/database"
)

type userRoleRepository struct {
	db *database.GormInstance
}

func NewUserRoleRepository(db *database.GormInstance) contract.UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (u *userRoleRepository) AssignUserRoles(ctx context.Context, userRoles []*models.UserRoleParam) error {
	return u.db.GetConn(ctx).Create(userRoles).Error
}
