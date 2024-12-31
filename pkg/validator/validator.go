package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	goValidator "github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gitub.com/umardev500/gopos/pkg/model"
	"gitub.com/umardev500/gopos/pkg/util"
)

type Validator interface {
	// Struct validates a struct
	Struct(s interface{}) Validator

	// Response returns the validation errors with response
	Response() *util.Response

	// Items returns all validation errors
	Items() []model.ValidationErr
}

type validator struct {
	validate *goValidator.Validate
	errs     []model.ValidationErr
	mu       sync.Mutex
}

var (
	validatorInstance *validator
	validatorOnce     sync.Once
)

func NewValidator() Validator {
	validatorOnce.Do(func() {
		validatorInstance = &validator{
			validate: goValidator.New(),
		}
	})

	return validatorInstance
}

func (v *validator) Struct(s interface{}) Validator {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.errs = nil // Clear previous errors

	// Validate the struct
	err := v.validate.Struct(s)
	if err == nil {
		return validatorInstance
	}

	value := reflect.ValueOf(s)

	// Check if the struct is a pointer
	if value.Kind() != reflect.Ptr {
		log.Fatal().Err(err).Msg("Struct must be a pointer")
	}

	// Check if the struct is a pointer to a struct
	if value.Kind() == reflect.Ptr && value.Elem().Kind() != reflect.Struct {
		log.Fatal().Err(err).Msg("Struct must be a pointer to a struct")
	}

	for _, fe := range err.(goValidator.ValidationErrors) {
		fn := fe.StructField()
		f, ok := value.Elem().Type().FieldByName(fn)
		if !ok {
			log.Fatal().Err(fmt.Errorf("field %s not found", fn)).Msg("Field not found")
		}

		item := v.parseErr(fe, f)
		v.errs = append(v.errs, *item)
	}

	return v
}

func (v *validator) parseErr(fe goValidator.FieldError, field reflect.StructField) *model.ValidationErr {
	tag := fe.Tag()
	jsonTag := field.Tag.Get("json")

	result := model.ValidationErr{
		Tag:  tag,
		Kind: fe.Kind().String(),
		Path: jsonTag,
	}

	switch tag {
	case "required":
		result.Message = "This field is required"
	case "min":
		result.Message = "This field must be at least " + fe.Param()
	case "max":
		result.Message = "This field must be at most " + fe.Param()
	case "len":
		result.Message = "This field must be " + fe.Param() + " characters long"
	case "email":
		result.Message = "This field must be a valid email address"
	case "oneof":
		result.Options = strings.Split(fe.Param(), " ")
		result.Message = "This field must be one of the options"
	}

	return &result
}

func (v *validator) Response() *util.Response {
	v.mu.Lock()
	defer v.mu.Unlock()

	if len(v.errs) > 0 {
		return util.ValidationResponse(nil, v.errs)
	}

	return nil
}

func (v *validator) Items() []model.ValidationErr {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.errs
}
