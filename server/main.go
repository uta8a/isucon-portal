package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://127.0.0.1:8080/api/auth/callback",
		Scopes:       []string{}, // 最小限。public dataのみを取得
		Endpoint:     github.Endpoint,
	}
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	var authCode string
	fmt.Scan(&authCode)
	fmt.Printf("%v\n", authCode)
	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext, tok)
	res, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}
