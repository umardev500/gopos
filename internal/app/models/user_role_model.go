package models

type UserRole struct {
	UserID string
	RoleID string
}

func (UserRole) TableName() string {
	return "user_roles"
}
