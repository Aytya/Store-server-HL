package domain

import "time"

type User struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	RegistrationDate time.Time `json:"registration_date"`
	Role             string    `json:"role"`
}

var UserBaseMessages = map[string]string{
	"required": "is required",
	"email":    "is not valid",
	"oneof":    "must be either 'admin' or 'client'",
}
