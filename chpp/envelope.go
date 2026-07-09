// Package chpp defines the typed Go representation of every CHPP XML file:
// one envelope struct per file (e.g. ClubXML, MatchesXML), the shared types
// and enums used across them, and the generic decode/validation logic that
// turns raw XML bytes into those structs.
package chpp

import (
	"encoding/xml"
	"errors"
)

// Envelope holds the fields common to every CHPP response file. Every
// generated XXXXML type embeds it, which is what lets DecodeXML validate
// any of them generically.
type Envelope struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`
}

func (e Envelope) chppFileName() string { return e.FileName }
func (e Envelope) chppError() string    { return e.Error }

// XMLEnvelope is satisfied by every CHPP response type via its embedded
// Envelope. Exported so other packages (parsed) can name it as a type
// constraint for their own generic helpers; its methods stay unexported
// since there's no legitimate reason to call them directly.
type XMLEnvelope interface {
	chppFileName() string
	chppError() string
}

// DecodeXML unmarshals buf into a new T and validates it: a non-empty
// <Error> becomes a Go error, and <FileName> must equal apiFile+".xml"
// (guards against e.g. decoding a club response as a team response).
func DecodeXML[T XMLEnvelope](buf []byte, apiFile string) (*T, error) {
	var x *T
	if err := xml.Unmarshal(buf, &x); err != nil {
		return nil, err
	}

	e := *x
	if e.chppError() != "" {
		return nil, errors.New(e.chppError())
	}
	if e.chppFileName() != apiFile+".xml" {
		return nil, errors.New("failed to parse the right " + apiFile + " type")
	}

	return x, nil
}
