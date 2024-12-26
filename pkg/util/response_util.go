package util

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/pkg/constant"
	"gitub.com/umardev500/gopos/pkg/logger"
	"gitub.com/umardev500/gopos/pkg/model"
)

type Response struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Code       string      `json:"code,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
	Reference  string      `json:"reference,omitempty"`
}

// ValidationResponse returns a common validation error response
func ValidationResponse(err error, fields []model.ValidationErr) *Response {
	ref := logger.LogError(fmt.Errorf("validation error"))

	return &Response{
		StatusCode: fiber.ErrUnprocessableEntity.Code,
		Message:    "Validation error",
		Code:       constant.ValidationErrorCodeName,
		Errors:     fields,
		Reference:  ref,
	}
}
