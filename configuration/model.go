package configuration

import "github.com/dgrijalva/jwt-go"

// Configuration main configuration struct
type Configuration struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Address     string   `json:"address"`
	SecretKey   string   `json:"jwt_secret_key"`
	Database    database `json:"database"`
}

type database struct {
	Driver string `json:"driver"`
	Url    string `json:"url"`
}

type Session struct {
	Name       *string    `json:"name,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Permission *UserLevel `json:"permission,omitempty"`
	jwt.StandardClaims
}

// UserLevel permission model of user
type UserLevel struct {
	IsAdmin *bool   `json:"is_admin,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
}
