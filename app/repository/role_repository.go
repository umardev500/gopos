package repository

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/contract"
	"gitub.com/umardev500/gopos/internal/app/models"
	"gitub.com/umardev500/gopos/pkg/database"
)

type roleRepository struct {
	db *database.GormInstance
}

func NewRoleRepository(db *database.GormInstance) contract.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) CountRolesByTenantID(ctx context.Context, roles []string, tenantID *string) (int64, error) {
	conn := r.db.GetConn(ctx)
	var count int64
	query := conn.Model(&models.Role{})

	// if tenant provided then it is tenant scoped
	// eles it's platform scoped
	if tenantID != nil {
		query.Where("tenant_id = ?", *tenantID)
	} else {
		query.Where("tenant_id IS NULL")
	}

	result := query.
		Where("id IN ?", roles).
		Count(&count)

	return count, result.Error
}
