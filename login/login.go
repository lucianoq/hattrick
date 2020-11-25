package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mrjones/oauth"
)

func main() {
	c := oauth.NewConsumer(
		prompt("Insert Consumer Key:"),
		prompt("Insert Consumer Secret:"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://chpp.hattrick.org/oauth/request_token.ashx",
			AuthorizeTokenUrl: "https://chpp.hattrick.org/oauth/authorize.aspx",
			AccessTokenUrl:    "https://chpp.hattrick.org/oauth/access_token.ashx",
		})

	requestToken, u, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	scopes := make([]string, 0)

	if yesNo("Do you want to be able to manage challenges?") {
		scopes = append(scopes, "manage_challenges")
	}
	if yesNo("Do you want to be able to set match orders?") {
		scopes = append(scopes, "set_matchorder")
	}
	if yesNo("Do you want to be able to manage youth players?") {
		scopes = append(scopes, "manage_youthplayers")
	}
	if yesNo("Do you want to be able to set training?") {
		scopes = append(scopes, "set_training")
	}
	if yesNo("Do you want to be able to place bids?") {
		scopes = append(scopes, "place_bid")
	}

	loginURL := u

	if len(scopes) > 0 {
		loginURL += "&scope=" + strings.Join(scopes, ",")
	}

	fmt.Println()
	fmt.Println("1. Go to --> " + loginURL)
	fmt.Println("2. Grant access, you should get back a verification code")

	verificationCode := prompt("3. Paste verification code here:")

	accToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("AccessToken = %s\n", accToken.Token)
	fmt.Printf("AccessSecret = %s\n", accToken.Secret)
	fmt.Printf("Access Additional Data = %+v\n", accToken.AdditionalData)
}

func prompt(question string) string {
	for {
		fmt.Print(question + " ")

		resp := ""
		_, err := fmt.Scanln(&resp)
		if err == nil {
			return strings.TrimSpace(resp)
		}
	}
}

func yesNo(question string) bool {
	for {
		fmt.Print(question + " [y/n] ")

		resp := ""
		_, err := fmt.Scanln(&resp)
		if err != nil {
			continue
		}

		switch strings.ToLower(resp) {

		case "y", "yes":
			return true

		case "n", "no":
			return false

		default:
			continue
		}
	}
}
