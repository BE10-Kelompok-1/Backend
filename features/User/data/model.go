package data

import (
	// "backend/domain"

	"backend/domain"
	"backend/features/Post/data"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname    string      `json:"firstname" validate:"required"`
	Lastname     string      `json:"lastname" validate:"required"`
	Username     string      `json:"username" validate:"required"`
	Email        string      `json:"email" validate:"required,email"`
	Password     string      `json:"password" validate:"required"`
	Birthdate    string      `json:"birthdate" validate:"required"`
	Photoprofile string      `json:"photoprofile"`
	Posts        []data.Post `gorm:"foreignKey:Userid"`
}

type User_Posting struct {
	Username  string
	Postid    uint
	Photo     string
	Caption   string
	CreatedAt time.Time
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:           int(u.ID),
		Firstname:    u.Firstname,
		Lastname:     u.Lastname,
		Username:     u.Username,
		Email:        u.Email,
		Password:     u.Password,
		Birthdate:    u.Birthdate,
		Photoprofile: u.Photoprofile,
	}
}

func (up *User_Posting) ToModelUserPosting() domain.User_Posting {
	return domain.User_Posting{
		Username:  up.Username,
		Postid:    uint(up.Postid),
		Photo:     up.Photo,
		Caption:   up.Caption,
		CreatedAt: up.CreatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func ParseToArr2(arr []User_Posting) []domain.User_Posting {
	var res []domain.User_Posting

	for _, val := range arr {
		res = append(res, val.ToModelUserPosting())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.ID = uint(data.ID)
	res.Firstname = data.Firstname
	res.Lastname = data.Lastname
	res.Username = data.Username
	res.Email = data.Email
	res.Password = data.Password
	res.Birthdate = data.Birthdate
	res.Photoprofile = data.Photoprofile

	return res
}
