package auth

import "github.com/isaqueveras/authentication-microservice/configuration/database"

// Service models a service base struct
type Service struct {
	repo IAuth
}

// GetService retrieves a service type
func GetService(r IAuth) *Service {
	return &Service{repo: r}
}

// GetAuthRepository retrieve repository for access to auth data
func GetAuthRepository(db *database.DBTransaction) IAuth {
	return New(db)
}
