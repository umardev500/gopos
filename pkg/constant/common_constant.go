package constant

type ContextKey int

const (
	TransactionContextKey ContextKey = iota
	ScopeContextKey
)

var (
	ValidationErrorCodeName string = "VALIDATION_ERROR"
)
