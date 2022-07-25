package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID           int
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	Password     string
	Birthdate    time.Time
	Photoprofile string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Search() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserUseCase interface {
	Register(newuser User, cost int) (User, error)
	UpdateUser(newuser User, userid, cost int) (User, error)
	SearchUser(username string) (User, error)
	DeleteUser(userid int) (bool, error)
}

type UserData interface {
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	SearchUserData(userid int) User
	DeleteUserData(userid int) bool
}
