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
<<<<<<< HEAD
	Firstname    string                `json:"firstname" form:"firstname" validate:"required"`
	Lastname     string                `json:"lastname" form:"lastname" validate:"required"`
	Username     string                `json:"username" form:"username" validate:"required"`
	Email        string                `json:"email" form:"email" validate:"required,email"`
	Password     string                `json:"password" form:"password" validate:"required"`
	Birthdate    string                `json:"birthdate" form:"birthdate" validate:"required"`
	Photoprofile string                `json:"photoprofile" form:"photoprofile"`
	Posts        []postdata.Post       `gorm:"foreignKey:Userid"`
=======
	Firstname    string          `json:"firstname" form:"firstname" validate:"required"`
	Lastname     string          `json:"lastname" form:"lastname" validate:"required"`
	Username     string          `json:"username" form:"username" validate:"required"`
	Email        string          `json:"email" form:"email" validate:"required,email"`
	Password     string          `json:"password" form:"password" validate:"required"`
	Birthdate    string          `json:"birthdate" form:"birthdate" validate:"required"`
	Photoprofile string          `json:"photoprofile" form:"photoprofile"`
	Posts        []postdata.Post `gorm:"foreignKey:Userid"`
>>>>>>> 13dfa70 (Update model.go)
	Comments     []commentdata.Comment `gorm:"foreignKey:Userid"`
}

type UserPosting struct {
	ID        int
	Photo     string
	Caption   string
	CreatedAt time.Time
}

type CommentUser struct {
	Id           int
	Firstname    string
	Lastname     string
	Photoprofile string
	Postid       int
	Comment      string
	Created_at   time.Time
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
		ID:        up.ID,
		Photo:     up.Photo,
		Caption:   up.Caption,
		CreatedAt: up.CreatedAt,
	}
}

func (cu *CommentUser) ToCommentUser() domain.CommentUser {
	return domain.CommentUser{
		Id:           cu.Id,
		Firstname:    cu.Firstname,
		Lastname:     cu.Lastname,
		Photoprofile: cu.Photoprofile,
		Postid:       cu.Postid,
		Comment:      cu.Comment,
		Created_at:   cu.Created_at,
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

func ParseCommentUserToArr(arr []CommentUser) []domain.CommentUser {
	var res []domain.CommentUser

	for _, val := range arr {
		res = append(res, val.ToCommentUser())
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
