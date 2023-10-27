// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc/auth/v1/user.proto

package authv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/uta8a/isucon-portal/server/gen/rpc/auth/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_7_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "rpc.auth.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/rpc.auth.v1.UserService/GetUser"
	// UserServiceLoginProcedure is the fully-qualified name of the UserService's Login RPC.
	UserServiceLoginProcedure = "/rpc.auth.v1.UserService/Login"
	// UserServiceLogoutProcedure is the fully-qualified name of the UserService's Logout RPC.
	UserServiceLogoutProcedure = "/rpc.auth.v1.UserService/Logout"
	// UserServiceCallbackProcedure is the fully-qualified name of the UserService's Callback RPC.
	UserServiceCallbackProcedure = "/rpc.auth.v1.UserService/Callback"
)

// UserServiceClient is a client for the rpc.auth.v1.UserService service.
type UserServiceClient interface {
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
	Callback(context.Context, *connect.Request[v1.CallbackRequest]) (*connect.Response[v1.CallbackResponse], error)
}

// NewUserServiceClient constructs a client for the rpc.auth.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUser: connect.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			opts...,
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+UserServiceLoginProcedure,
			opts...,
		),
		logout: connect.NewClient[v1.LogoutRequest, v1.LogoutResponse](
			httpClient,
			baseURL+UserServiceLogoutProcedure,
			opts...,
		),
		callback: connect.NewClient[v1.CallbackRequest, v1.CallbackResponse](
			httpClient,
			baseURL+UserServiceCallbackProcedure,
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUser  *connect.Client[v1.GetUserRequest, v1.GetUserResponse]
	login    *connect.Client[v1.LoginRequest, v1.LoginResponse]
	logout   *connect.Client[v1.LogoutRequest, v1.LogoutResponse]
	callback *connect.Client[v1.CallbackRequest, v1.CallbackResponse]
}

// GetUser calls rpc.auth.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// Login calls rpc.auth.v1.UserService.Login.
func (c *userServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Logout calls rpc.auth.v1.UserService.Logout.
func (c *userServiceClient) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return c.logout.CallUnary(ctx, req)
}

// Callback calls rpc.auth.v1.UserService.Callback.
func (c *userServiceClient) Callback(ctx context.Context, req *connect.Request[v1.CallbackRequest]) (*connect.Response[v1.CallbackResponse], error) {
	return c.callback.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the rpc.auth.v1.UserService service.
type UserServiceHandler interface {
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
	Callback(context.Context, *connect.Request[v1.CallbackRequest]) (*connect.Response[v1.CallbackResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceGetUserHandler := connect.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		opts...,
	)
	userServiceLoginHandler := connect.NewUnaryHandler(
		UserServiceLoginProcedure,
		svc.Login,
		opts...,
	)
	userServiceLogoutHandler := connect.NewUnaryHandler(
		UserServiceLogoutProcedure,
		svc.Logout,
		opts...,
	)
	userServiceCallbackHandler := connect.NewUnaryHandler(
		UserServiceCallbackProcedure,
		svc.Callback,
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/rpc.auth.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceLoginProcedure:
			userServiceLoginHandler.ServeHTTP(w, r)
		case UserServiceLogoutProcedure:
			userServiceLogoutHandler.ServeHTTP(w, r)
		case UserServiceCallbackProcedure:
			userServiceCallbackHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.auth.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.auth.v1.UserService.Login is not implemented"))
}

func (UnimplementedUserServiceHandler) Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.auth.v1.UserService.Logout is not implemented"))
}

func (UnimplementedUserServiceHandler) Callback(context.Context, *connect.Request[v1.CallbackRequest]) (*connect.Response[v1.CallbackResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("rpc.auth.v1.UserService.Callback is not implemented"))
}
