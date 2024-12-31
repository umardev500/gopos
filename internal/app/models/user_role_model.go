package models

type UserRole struct{}

func (UserRole) TableName() string {
	return "user_roles"
}

type UserRoleParam struct {
	UserID string
	RoleID string
}

func (UserRoleParam) TableName() string {
	return "user_roles"
}
