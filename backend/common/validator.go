package common

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mnabila/mytask/internal/entities"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"

	translations "github.com/go-playground/validator/v10/translations/id"
)

// ValidateRequest performs validation on the provided data using the struct tags and validation rules.
// It takes a single parameter: `data` (any) representing the data to be validated.
// The function returns a slice of `entities.FieldError` containing any validation errors encountered during the validation process.
func ValidateRequest(data any) []entities.FieldError {
	// Create a new language translator for the 'id' locale
	lang := id.New()
	uni := ut.New(lang, lang)

	// Get the translator for the 'id' locale
	trans, _ := uni.GetTranslator("id")

	// Create a new validator instance
	validate := validator.New()

	// Register default translations for the validator
	translations.RegisterDefaultTranslations(validate, trans)

	// Register a custom tag name function for the validator
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Perform the data validation
	err := validate.Struct(data)

	// Check if there are validation errors
	if errors.As(err, &validator.ValidationErrors{}) {
		var outs []entities.FieldError
		for _, e := range err.(validator.ValidationErrors) {
			// Create a new FieldError instance for each validation error encountered
			outs = append(outs, entities.FieldError{
				Field:   e.Field(),
				Message: e.Translate(trans),
			})
		}
		return outs
	}

	// Return nil if there are no validation errors
	return nil
}
