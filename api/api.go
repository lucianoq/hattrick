package api

import "github.com/lucianoq/hattrick/parsed"

// NewAPI ...
func NewAPI(
	consumerKey, consumerSecret,
	accessToken, accessSecret string,
	additionalData map[string]string,
) (*API, error) {

	p, err := parsed.NewParsed(consumerKey, consumerSecret, accessToken, accessSecret, additionalData)
	if err != nil {
		return nil, err
	}
	return &API{
		parsed: p,
	}, nil
}

// API ...
type API struct {
	parsed parsed.Parsed
}
