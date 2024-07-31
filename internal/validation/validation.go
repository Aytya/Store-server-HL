package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func HandleValidationErrors(validationErrors validator.ValidationErrors, baseMessages map[string]string) string {
	var errorMessages []string
	for _, err := range validationErrors {
		if fieldMessage, found := baseMessages[err.Tag()]; found {
			errorMessages = append(errorMessages, err.Field()+" "+fieldMessage)
		} else {
			errorMessages = append(errorMessages, "Invalid field: "+err.Field())
		}
	}
	return strings.Join(errorMessages, ", ")
}
