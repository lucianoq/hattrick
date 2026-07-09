package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	WorldLanguagesAPIFile    = "worldlanguages"
	WorldLanguagesAPIVersion = "1.2"
)

// WorldLanguagesXML contains the list of all languages supported by
// Hattrick.
type WorldLanguagesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The list of all languages supported by Hattrick.
	Languages []*WorldLanguage `xml:"LanguageList>Language"`
}

// WorldLanguage is a single language supported by Hattrick.
type WorldLanguage struct {
	ID   id.Language `xml:"LanguageID"`
	Name string      `xml:"LanguageName"`
}
