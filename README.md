# Hattrick API

[![License](https://img.shields.io/:license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/lucianoq/hattrick?status.svg)](https://godoc.org/github.com/lucianoq/hattrick)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucianoq/hattrick)](https://goreportcard.com/report/github.com/lucianoq/hattrick)

Hattrick CHPP API for Go.

This package provides a list of Go types and functions to access online game 
Hattrick (www.hattrick.org) data, using and abstracting the official CHPP 
interface.

If you plan to write a Go application, this package helps you to obtain the data
with strong typing and types/constants definition, just ignoring the CHPP XML 
files, versions and messy unsigned integer list of constants.


## Example

```go
package main

import (
	"fmt"

	"github.com/lucianoq/hattrick"
)

func main() {
	ht, err := hattrick.NewClient(
		hattrick.AuthConfig{
			ConsumerKey:          "see authentication section",
			ConsumerSecret:       "see authentication section",
			AccessToken:          "see authentication section",
			AccessSecret:         "see authentication section",
		},
	)
	// handle err

	players, err := ht.GetMyPlayers()
	// handle err

	for _, p := range players {
		fmt.Printf("%d %s %s\n", p.PlayerNumber, p.FirstName, p.LastName)
	}
}
```


## Contribution
I focused on the most important CHPP files, so we can, right now, easily extract 
almost all player/match/team related data, but there is still a lot of work to 
do on, e.g., youth, transfers, national teams.

If you want to contribute to the project, just drop me an email (on github profile).


## Authentication

( See www.hattrick.org/Community/CHPP/oauth/ for more details. )

The Hattrick Client requires 4 strings.

Two about the CHPP Application (or Consumer for oauth):
 - `ConsumerKey`
 - `ConsumerSecret`
 
Two about the authorised user (or Resource Owner for oauth):
 - `AccessToken`
 - `AccessSecret`
 
##### Consumer

You can obtain the first two creating a new CHPP application on the Hattrick 
website. Request must be motivated and approved by Hattrick developers.

##### Resource owner

You need to generate a pair of Access string, token and secret, for every user 
that needs to access Hattrick information.

For simplicity, the process has been coded in `login` package.
You can just run it to obtain the resource owner strings.