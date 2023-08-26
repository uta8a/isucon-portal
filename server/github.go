package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"connectrpc.com/connect"
	authv1 "github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type User = authv1.User

type GitHubConf struct {
	baseUrl string
	conf    *oauth2.Config
}

func (g *GitHubConf) New() {
	g.baseUrl = os.Getenv("BASE_URL")
	g.conf = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  fmt.Sprintf("%s/api/auth/callback", g.baseUrl),
		Scopes:       []string{}, // 最小限。public dataのみを取得
		Endpoint:     github.Endpoint,
	}
}

func (g *GitHubConf) IssueURL(state string) string {
	url := g.conf.AuthCodeURL(state)
	return url
}

func (g *GitHubConf) GetUser(authCode string) User {
	// Handle the exchange code to initiate a transport.
	tok, err := g.conf.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatal(err)
	}
	client := g.conf.Client(context.TODO(), tok)
	res, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var u User
	if err = json.Unmarshal(body, &u); err != nil {
		log.Fatal(err)
	}
	return u
}

const tokenHeader = "Acme-Token"

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// Send a token with client requests.
				req.Header().Set(tokenHeader, "sample")
			} else if req.Header().Get(tokenHeader) == "" {
				// Check token in handlers.
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
