package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	authv1 "github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1"
	"github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1/authv1connect"
)

func TestGetUser(t *testing.T) {
	t.Parallel()
	mux := App()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		userId   string
		want     string
	}{
		{
			scenario: "ユーザーが存在する場合",
			userId:   "test",
			want:     "test",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.userId, func(t *testing.T) {
			t.Parallel()
			client := authv1connect.NewUserServiceClient(
				server.Client(),
				server.URL,
			)
			res, err := client.GetUser(
				context.Background(),
				connect.NewRequest(&authv1.GetUserRequest{
					UserId: c.userId,
				}),
			)
			if err != nil {
				t.Error(err)
			}
			if res.Msg.GetUser().UserId != c.want {
				t.Errorf("auth want %s, got %s", c.want, res.Msg.GetUser().UserId)
			}
		})
	}
}
