package postgres

import (
	"github.com/isaqueveras/authentication-microservice/configuration/database"
	infra "github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser"
)

// PGAuth implements methods for postgres query execution
type PGAuth struct {
	DB *database.DBTransaction
}

// CreateUser create user on database
func (pg *PGAuth) CreateUser(in infra.User) (err error) {
	if err = pg.DB.Builder.
		Insert("public.t_users").
		Columns("name", "email", "passw").
		Values(in.Name, in.Email, in.Passw).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return err
	}

	return nil
}
