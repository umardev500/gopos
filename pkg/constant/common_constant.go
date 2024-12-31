package constant

import "fmt"

type ContextKey int

const (
	TransactionContextKey ContextKey = iota
	ScopeContextKey
	ClaimsContextKey
)

var (
	ValidationErrorCodeName   string = "VALIDATION_ERROR"
	ConflictErrorCodeName     string = "CONFLICT_ERROR"
	BadRequestErrorCodeName   string = "BAD_REQUEST_ERROR"
	InternalErrorCodeName     string = "INTERNAL_ERROR"
	RoleCountNotValidCodeName string = "ROLE_COUNT_NOT_VALID_ERR"
)

var (
	ErrInvalidToken      error = fmt.Errorf("invalid token")
	ErrRoleCountNotValid error = fmt.Errorf("role count not valid")
)
