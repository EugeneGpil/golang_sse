package response

import (
	"github.com/EugeneGpil/golang_sse/app/ship/translator"

	responsePackage "github.com/EugeneGpil/response"
)

func WriteValidationErrors(response responsePackage.Response, Code int, Errors map[string]string) {
	response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Message: translator.Translate("errors.request_contains_validation_errors"),
		Errors:  Errors,
		Code:    Code,
	})
}
