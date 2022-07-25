package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
	Birthdate time.Time
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type UserUseCase interface {
	Register(newuser User, cost int) (User, error)
	UpdateUser(newuser User, userid, cost int) (User, error)
}

type UserData interface {
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
}
