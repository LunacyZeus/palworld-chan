package custom

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type XValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type GlobalErrorHandlerResp struct {
	Code    int    `json:"code"`
	Result  any    `json:"result"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

var validate = validator.New()
var instance *XValidator
var once sync.Once

func CustomValidator() *XValidator {
	once.Do(func() {
		GlobalValidator := &XValidator{
			validator: validate,
		}
		instance = GlobalValidator
	})
	return instance
}
