package translator

import (
	"github.com/EugeneGpil/golang_sse/app/ship/translator/lang"

	translatorPackage "github.com/EugeneGpil/translator"
)

var fallbackLang = "En"
var translations = map[string]map[string]string{
	"En": lang.En,
	"Ru": lang.Ru,
}
var translator translatorPackage.Translator

func Init() {
	translator = translatorPackage.New()

	translator.SetFallbackLang(fallbackLang)

	translator.SetTranslations(translations)
}

func Translate(key string) string {
	return translator.
		SetKey(key).
		Run()
}
