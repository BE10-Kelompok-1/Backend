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
	Birthdate    string
	Photoprofile string
	Posts        []Post
}

type User_Posting struct {
	Username  string
	Postid    uint
	Photo     string
	Caption   string
	CreatedAt time.Time
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Search() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User, cost int) int
	UpdateUser(newuser User, userid, cost int) int
	SearchUser(username string) (User, []User_Posting, int)
	DeleteUser(userid int) int
	LoginUser(userdata User) (User, error)
}

type UserData interface {
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	SearchUserData(username string) (User, []User_Posting)
	DeleteUserData(userid int) bool
	LoginData(userdata User) User
	GetPasswordData(name string) string
}
