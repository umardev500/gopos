package contract

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/models"
)

type UserTenantRepository interface {
	// Assing user to tenant
	AssignUserToTenant(ctx context.Context, userTenant *models.UserTenant) error
}
