package delivery

import (
	"backend/domain"
)

type UserFormat struct {
	Firstname    string `json:"firstname" validate:"required"`
	Lastname     string `json:"lastname" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Birthdate    string `json:"birthdate" validate:"required"`
	Photoprofile string `json:"photoprofile"`
}

func (i *UserFormat) ToModel() domain.User {
	return domain.User{
		Firstname:    i.Firstname,
		Lastname:     i.Lastname,
		Username:     i.Username,
		Email:        i.Email,
		Password:     i.Password,
		Birthdate:    i.Birthdate,
		Photoprofile: i.Photoprofile,
	}
}
