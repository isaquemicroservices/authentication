package postgres

import "github.com/isaqueveras/authentication-microservice/configuration/database"

// PGAuth implements methods for postgres query execution
type PGAuth struct {
	DB *database.DBTransaction
}
