package authuser

import "github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser"

// IAuth defines all services available for authentication
type IAuth interface {
	CreateUser(user authuser.User) error
	GetUser(email *string) (*authuser.User, error)
}
