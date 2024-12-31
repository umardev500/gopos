package repository

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/database"
)

type userTenantRepository struct {
	db *database.GormInstance
}

func NewUserTenantRepository(db *database.GormInstance) contract.UserTenantRepository {
	return &userTenantRepository{
		db: db,
	}
}

func (r *userTenantRepository) AssignUserToTenant(ctx context.Context, userTenant *models.UserTenant) error {
	conn := r.db.GetConn(ctx)

	return conn.Create(&userTenant).Error
}
