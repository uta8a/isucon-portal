package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	authv1 "github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1"
	"github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AppServer struct{}

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

func App() http.Handler {
	appServer := &AppServer{}
	mux := http.NewServeMux()
	path, handler := authv1connect.NewUserServiceHandler(appServer)
	mux.Handle(path, handler)
	return mux
}

// register handlers
func Run() {
	mux := App()
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
