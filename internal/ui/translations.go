package ui

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"rompebolas/data"
	"strings"
)

var Loc *i18n.Localizer

func SetLanguage(lang string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.MustParseMessageFileBytes(data.GetAssetBytes("translations/active.es.json"), "translations/active.es.json")

	if lang == "" {
		lang = os.Getenv("LANG")
		if i := strings.Index(lang, ".UTF-8"); i != -1 {
			lang = lang[:i]
		}
	}

	Loc = i18n.NewLocalizer(bundle, lang)
}
