package data

import (
	"backend/domain"
	commentdata "backend/features/Comment/data"
	postdata "backend/features/Post/data"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname    string                `json:"firstname" form:"firstname" validate:"required"`
	Lastname     string                `json:"lastname" form:"lastname" validate:"required"`
	Username     string                `json:"username" form:"username" validate:"required"`
	Email        string                `json:"email" form:"email" validate:"required,email"`
	Password     string                `json:"password" form:"password" validate:"required"`
	Birthdate    string                `json:"birthdate" form:"birthdate" validate:"required"`
	Photoprofile string                `json:"photoprofile" form:"photoprofile"`
	Posts        []postdata.Post       `gorm:"foreignKey:Userid"`
	Comments     []commentdata.Comment `gorm:"foreignKey:Userid"`
}

type UserPosting struct {
	UserID    uint
	PostID    int
	Photo     string
	Caption   string
	CreatedAt time.Time
	Posts     []postdata.Post `gorm:"foreignKey:Userid"`
}

type UserPostingComment struct {
	PostID       int
	CommentID    int
	Firstname    string
	Lastname     string
	Photoprofile string
	Comment      string
	CreatedAt    time.Time
	Comments     []commentdata.Comment `gorm:"foreignKey:Userid"`
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

func (up *UserPosting) ToUserPosting() domain.UserPosting {
	return domain.UserPosting{
		PostID:    int(up.PostID),
		Photo:     up.Photo,
		Caption:   up.Caption,
		CreatedAt: up.CreatedAt,
	}
}

func (upc *UserPostingComment) ToUserPostingComment() domain.UserPostingComment {
	return domain.UserPostingComment{
		CommentID:    int(upc.CommentID),
		Firstname:    upc.Firstname,
		Lastname:     upc.Lastname,
		Photoprofile: upc.Photoprofile,
		Comment:      upc.Comment,
		CreatedAt:    upc.CreatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func ParseUserPostingToArr(arr []UserPosting) []domain.UserPosting {
	var res []domain.UserPosting

	for _, val := range arr {
		res = append(res, val.ToUserPosting())
	}

	return res
}

func ParseUserPostingCommentToArr(arr []UserPostingComment) []domain.UserPostingComment {
	var res []domain.UserPostingComment

	for _, val := range arr {
		res = append(res, val.ToUserPostingComment())
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
