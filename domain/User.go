package domain

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID           int
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	Password     string
	Birthdate    string
	Photoprofile string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Search() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User, cost int) int
	UpdateUser(newuser User, userid, cost int) int
	SearchUser(username string) (User, error)
	DeleteUser(userid int) (bool, error)
}

type UserData interface {
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	SearchUserData(username string) (User, error)
	DeleteUserData(userid int) bool
}
