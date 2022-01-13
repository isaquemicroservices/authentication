package authuser

import (
	"github.com/isaqueveras/authentication-microservice/configuration/database"
	"github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser"
	"github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser/postgres"
)

// repository is a base structure that implements methods specified by IAuth
type repository struct {
	pg *postgres.PGAuth
}

// New creates a new user repository
func New(db *database.DBTransaction) *repository {
	return &repository{pg: &postgres.PGAuth{DB: db}}
}

func (r *repository) CreateUser(in authuser.User) error {
	return r.pg.CreateUser(in)
}

func (r *repository) GetUser(id *int64) (*authuser.User, error) {
	return r.pg.GetUser(id)
}
