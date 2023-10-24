package response

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/translator"

	responsePackage "github.com/EugeneGpil/response"
)

func WriteValidationErrors(writer http.ResponseWriter, Code int, Errors map[string]string) {
	response := responsePackage.New(writer)

	response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Message: translator.Translate("errors.request_contains_validation_errors"),
		Errors: Errors,
		Code: Code,
	})
}
