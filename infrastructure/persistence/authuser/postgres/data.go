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
		Columns("name", "email", "passw", "level_id").
		Values(in.Name, in.Email, in.Passw, in.Permission.ID).
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
		Select(`
			TU.id AS user_id, 
			TU.name AS user_name, 
			TU.email AS user_email, 
			TU.passw AS user_passw,
			TUL.id AS user_level_id,
			TUL.name AS user_level_name`,
		).
		From("public.t_users TU").
		Join("public.t_users_level TUL ON TUL.id = TU.level_id").
		Where(squirrel.Eq{
			"email": email,
		}).
		Limit(1).
		Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.Passw,
			&res.Permission.ID,
			&res.Permission.Name,
		); err != nil {
		return res, err
	}

	// Check user level
	if res.Permission.ID == infra.Permission["admin"] {
		res.Permission.IsAdmin = true
	}

	return res, nil
}
