package authuser

import (
	"context"

	"github.com/isaqueveras/authentication-microservice/configuration/database"
	domain "github.com/isaqueveras/authentication-microservice/domain/authuser"
	infra "github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser create user
func CreateUser(ctx context.Context, in *User) (err error) {
	var (
		pass        []byte
		transaction *database.DBTransaction
	)

	// Generate password
	if pass, err = bcrypt.GenerateFromPassword([]byte(in.Passw), 14); err != nil {
		return err
	}

	var data = infra.User{
		Name:  in.Name,
		Email: in.Email,
		Passw: string(pass),
	}

	// Initializing transaction with database
	if transaction, err = database.OpenConnection(ctx, false); err != nil {
		return err
	}

	defer transaction.Rollback()

	var repo = domain.New(transaction)
	if err = repo.CreateUser(data); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
