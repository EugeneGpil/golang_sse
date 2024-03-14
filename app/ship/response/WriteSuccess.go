package response

import (
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
	responsePackage "github.com/EugeneGpil/response"
)

func WriteSuccess(response responsePackage.Response) {
	response.WriteSuccess(translator.Translate("Success"))
}
