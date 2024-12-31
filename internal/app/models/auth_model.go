package models

type AuthUser struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password_hash"`
	TenantID     *string `json:"tenant_id,omitempty"` // From relation to user_tenants
	BranchID     *string `json:"branch_id,omitempty"` // From relation to user_branches
}

func (AuthUser) TableName() string {
	return "users"
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}
