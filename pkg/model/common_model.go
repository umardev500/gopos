package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// IdsRequest is a request for a list of ids
type IdsRequest struct {
	Ids []uuid.UUID `json:"ids" validate:"required,min=1"`
}

// StringSlice converts []uuid.UUID to []string
func (i IdsRequest) StringSlice() []string {
	var ids []string
	for _, id := range i.Ids {
		ids = append(ids, id.String())
	}
	return ids
}

type PaginatedResult struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// Parse pagination params to default if invalid
func (p PaginationParams) Parse() *PaginationParams {
	result := PaginationParams{
		Page: p.Page,
	}

	if p.Limit <= 0 {
		result.Limit = 10
	}

	return &result
}

// Time is a extension struct for created_at, updated_at, deleted_at
type Time struct {
	CreatedAt time.Time       `json:"created_at"`
	UpdateAt  *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// ValidationErr is a struct for validation error item
type ValidationErr struct {
	Tag     string      `json:"tag,omitempty"`
	Kind    string      `json:"kind,omitempty"`
	Path    string      `json:"path,omitempty"`
	Options interface{} `json:"options,omitempty"`
	Message string      `json:"message,omitempty"`
}
