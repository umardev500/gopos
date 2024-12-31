package contract

import "context"

type RoleRepository interface {
	CountRolesByTenantID(ctx context.Context, roles []string, tenantID *string) (int64, error)
}
