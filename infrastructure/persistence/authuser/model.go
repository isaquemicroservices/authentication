package authuser

import "time"

// User struct for user
type User struct {
	Id        int64
	Name      string
	Email     string
	Passw     string
	CreateAt  time.Time
	UpdatedAt time.Time
}
