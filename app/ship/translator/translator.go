package translator

import (
	"github.com/EugeneGpil/golang_sse/app/ship/translator/lang"
	
	"github.com/EugeneGpil/translator"
)

var translations map[string]map[string]string
var fallbackLang = "En"

func Init() {
	translator.SetFallbackLang(fallbackLang)

	translations := map[string]map[string]string{
		"En": lang.En,
		"Ru": lang.Ru,
	}

	translator.SetTranslations(translations)
}

func Translate(key string) string {
	return translator.Translate(key)
}
