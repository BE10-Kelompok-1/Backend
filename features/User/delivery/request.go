package delivery

import (
	"backend/domain"
)

type UserFormat struct {
	Firstname    string `json:"firstname" form:"firstname" validate:"required"`
	Lastname     string `json:"lastname" form:"lastname" validate:"required"`
	Username     string `json:"username" form:"username" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Password     string `json:"password" form:"password" validate:"required"`
	Birthdate    string `json:"birthdate" form:"birthdate" validate:"required"`
	Photoprofile string `json:"photoprofile" form:"photoprofile"`
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

type LoginFormat struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: i.Username,
		Password: i.Password,
	}
}
