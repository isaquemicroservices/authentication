package auth

import (
	"github.com/isaqueveras/authentication-microservice/configuration/database"
	"github.com/isaqueveras/authentication-microservice/infrastructure/persistence/auth/postgres"
)

// repository is a base structure that implements methods specified by IAuth
type repository struct {
	pg *postgres.PGAuth
}

// New creates a new user repository
func New(db *database.DBTransaction) *repository {
	return &repository{pg: &postgres.PGAuth{DB: db}}
}
