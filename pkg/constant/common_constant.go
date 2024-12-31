package constant

import "fmt"

type ContextKey int

const (
	TransactionContextKey ContextKey = iota
	ScopeContextKey
)

var (
	ValidationErrorCodeName string = "VALIDATION_ERROR"
)

var (
	ErrInvalidToken error = fmt.Errorf("invalid token")
)
