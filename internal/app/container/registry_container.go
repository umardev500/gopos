package container

import (
	"gitub.com/umardev500/gopos/pkg/contract"
	"gitub.com/umardev500/gopos/pkg/database"
	"gitub.com/umardev500/gopos/pkg/validator"
)

func NewRegistryContainer(db *database.GormInstance, v validator.Validator) []contract.Container {
	return []contract.Container{
		NewUserContainer(db, v),
	}
}
