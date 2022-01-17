package postgres

import (
	"github.com/Masterminds/squirrel"
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

// GetUser get data of user on database
func (pg *PGAuth) GetUser(email *string) (res *infra.User, err error) {
	res = new(infra.User)

	if err = pg.DB.Builder.
		Select("id, name, email, passw").
		From("public.t_users").
		Where(squirrel.Eq{"email": email}).
		Limit(1).
		Scan(&res.Id, &res.Name, &res.Email, &res.Passw); err != nil {
		return res, err
	}

	return res, nil
}
