package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorReasons := []rest_err.Reason{}

		for _, err := range validationErr.(validator.ValidationErrors) {
			reason := rest_err.NewReason(err.Field(), err.Translate(transl))
			errorReasons = append(errorReasons, reason)
		}

		return rest_err.NewBadRequestValidationError("some fields are invalid", errorReasons)
	} else {
		return rest_err.NewBadRequestError("error trying to convert fields")
	}
}
