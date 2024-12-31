package models

import "github.com/google/uuid"

type CreateUserTenantRequest struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	TenantID uuid.UUID `json:"tenant_id" validate:"required"`
}

func (CreateUserTenantRequest) TableName() string {
	return "user_tenants"
}
