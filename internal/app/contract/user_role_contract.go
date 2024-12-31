package contract

import (
	"context"

	"gitub.com/umardev500/gopos/internal/app/models"
)

type UserRoleRepository interface {
	AssignUserRoles(ctx context.Context, userRoles []*models.UserRoleParam) error
}
