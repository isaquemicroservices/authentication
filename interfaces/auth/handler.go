package auth

import app "github.com/isaqueveras/authentication-microservice/application/auth"

// Server implements proto interface
type Server struct {
	app.UnimplementedAuthServer
}
