package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	authv1 "github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1"
	"github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AppServer struct {
	g *GitHubConf
}

func (s *AppServer) GetUser(ctx context.Context, req *connect.Request[authv1.GetUserRequest]) (*connect.Response[authv1.GetUserResponse], error) {
	res := connect.NewResponse(&authv1.GetUserResponse{
		User: &authv1.User{
			UserId:      req.Msg.UserId,
			DisplayName: "テストユーザだよ",
			IconPath:    "https://avatars.example.com/u/12345678?v=4",
		},
	})
	return res, nil
}

func (s *AppServer) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	url := s.g.IssueURL("state") // CSRFされる可能性がないので固定
	res := connect.NewResponse(&authv1.LoginResponse{
		Url: url,
	})
	return res, nil
}

func (s *AppServer) Callback(ctx context.Context, req *connect.Request[authv1.CallbackRequest]) (*connect.Response[authv1.CallbackResponse], error) {
	fmt.Println(req.Msg.Code)
	res := connect.NewResponse(&authv1.CallbackResponse{
		User: &authv1.User{
			UserId:      "12345678",
			DisplayName: "テストユーザだよ",
			IconPath:    "https://avatars.example.com/u/12345678?v=4",
		},
	})
	return res, nil
}

func (s *AppServer) Logout(ctx context.Context, req *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.LogoutResponse], error) {
	res := connect.NewResponse(&authv1.LogoutResponse{})
	return res, nil
}

// register handlers
func App(interceptor connect.UnaryInterceptorFunc) http.Handler {
	g := &GitHubConf{}
	g.New()
	appServer := &AppServer{
		g: g,
	}
	mux := http.NewServeMux()
	// GetUser by userId
	mux.Handle(authv1connect.NewUserServiceHandler(appServer))
	return mux
}

func Run() {
	mux := App(NewAuthInterceptor())
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
