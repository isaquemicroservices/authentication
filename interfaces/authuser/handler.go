package authuser

import (
	"context"

	app "github.com/isaqueveras/authentication-microservice/application/authuser"
)

// Server implements proto interface
type Server struct {
	app.UnimplementedAuthServer
}

// CreateUser create user on database
func (s *Server) CreateUser(ctx context.Context, in *app.User) (*app.Empty, error) {
	if err := app.CreateUser(ctx, in); err != nil {
		return nil, err
	}

	return nil, nil
}

// Login
func (s *Server) Login(ctx context.Context, in *app.Credentials) (res *app.User, err error) {
	if res, err = app.Login(ctx, in); err != nil {
		return res, err
	}

	return res, nil
}
