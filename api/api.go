// Package api implements the public CHPP client (API): one exported method
// per operation, each wrapping a parsed.Parsed fetch-and-decode call and
// returning already-decoded Go types instead of raw XML envelopes.
package api

import "github.com/lucianoq/hattrick/parsed"

// NewAPI creates a new API client authenticated with the given OAuth 1.0a
// consumer key/secret and access token/secret. additionalData is passed
// through to the underlying HTTP client for any extra per-request query
// parameters required by the CHPP endpoints used.
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

// API is the public CHPP client: every exported method wraps one CHPP
// file/action, returning already-decoded Go types.
type API struct {
	parsed parsed.Parsed
}
