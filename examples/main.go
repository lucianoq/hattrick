package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lucianoq/hattrick"
)

func main() {
	ht, err := hattrick.NewClient(
		hattrick.AuthConfig{
			ConsumerKey:    ConsumerKey,
			ConsumerSecret: ConsumerSecret,
			AccessToken:    AccessToken,
			AccessSecret:   AccessSecret,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	getMyPlayers(ht)
	getMatchesArchive(ht)
	getMyLeague(ht)

}

func getMyLeague(ht hattrick.Client) {
	league, err := ht.GetMyLeague()
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range league.Teams {
		fmt.Printf("%d %s %s\n", t.Position, t.TeamID, t.TeamName)
	}
}

func getMatchesArchive(ht hattrick.Client) {
	matches, err := ht.GetMatchesArchive(
		544828,
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range matches {
		fmt.Printf("%d %s VS %s  (%d - %d)\n",
			m.MatchID,
			m.HomeTeam.Name, m.AwayTeam.Name,
			m.HomeGoals, m.AwayGoals,
		)
	}
}

func getMyPlayers(ht hattrick.Client) {
	players, err := ht.GetMyPlayers()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range players {
		fmt.Printf("%d %s %s\n", p.Number, p.FirstName, p.LastName)
	}
}
