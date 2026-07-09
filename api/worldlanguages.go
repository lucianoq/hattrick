package api

import "github.com/lucianoq/hattrick/chpp"

// GetWorldLanguages shows the list of all languages supported by Hattrick.
func (a *API) GetWorldLanguages() ([]*chpp.WorldLanguage, error) {
	res, err := a.parsed.GetWorldLanguagesXML(nil)
	if err != nil {
		return nil, err
	}

	return res.Languages, nil
}
