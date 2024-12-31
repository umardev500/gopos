package models

type Role struct{}

func (Role) TableName() string {
	return "roles"
}
