package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetTranslations shows the translated names of all CHPP enum values, in
// the given language.
func (a *API) GetTranslations(languageID id.Language) (*chpp.TranslationsXML, error) {
	values := map[string]string{
		"languageId": languageID.String(),
	}

	return a.parsed.GetTranslationsXML(values)
}
