package authuser

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	config "github.com/isaqueveras/authentication-microservice/configuration"
	"github.com/isaqueveras/authentication-microservice/configuration/database"
	domain "github.com/isaqueveras/authentication-microservice/domain/authuser"
	infra "github.com/isaqueveras/authentication-microservice/infrastructure/persistence/authuser"
	"github.com/isaqueveras/authentication-microservice/utils"
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
		Permission: infra.UserLevel{
			ID: in.Permission.Id,
		},
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

// Login
func Login(ctx context.Context, in *Credentials) (res *User, err error) {
	res = new(User)

	var transaction *database.DBTransaction
	// Initializing transaction with database
	if transaction, err = database.OpenConnection(ctx, true); err != nil {
		return nil, err
	}

	defer transaction.Rollback()

	var repo = domain.New(transaction)
	data, err := repo.GetUser(utils.GetPointerString(in.Email))
	if err != nil {
		return nil, errors.New("User not found: " + err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Passw), []byte(in.Passw)); err != nil {
		return nil, err
	}

	if res.Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, config.Session{
		Name: utils.GetPointerString(data.Name),
		Permission: &config.UserLevel{
			IsAdmin: utils.GetPointerBool(data.Permission.IsAdmin),
			ID:      utils.GetPointerInt64(data.Permission.ID),
			Name:    utils.GetPointerString(data.Permission.Name),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
			Issuer:    "isaqueveras.auth",
			NotBefore: time.Now().Unix(),
		},
	}).SignedString([]byte(config.Get().SecretKey)); err != nil {
		return nil, errors.New("Could not generate token")
	}

	res.Id = data.Id
	res.Name = data.Name
	res.Email = data.Email
	res.Passw = ""

	return
}
