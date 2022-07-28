package data

import (
	"backend/domain"
	"backend/features/Comment/data"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Userid   int
	Photo    string         `json:"photo" form:"photo"`
	Caption  string         `json:"caption" form:"caption" validate:"required"`
	Comments []data.Comment `gorm:"foreignKey:Postid"`
}

type PostComent struct {
	ID           int
	Firstname    string
	Lastname     string
	Username     string
	Photoprofile string
	Photo        string
	Caption      string
	CreatedAt    time.Time
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

func (p *Post) ToModel() domain.Post {
	return domain.Post{
		ID:        int(p.ID),
		Userid:    p.Userid,
		Photo:     p.Photo,
		Caption:   p.Caption,
		CreatedAt: p.CreatedAt,
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

func (pc *PostComent) ToPostComent() domain.PostComent {
	return domain.PostComent{
		ID:           pc.ID,
		Firstname:    pc.Firstname,
		Lastname:     pc.Lastname,
		Username:     pc.Username,
		Photoprofile: pc.Photoprofile,
		Photo:        pc.Photo,
		Caption:      pc.Caption,
		CreatedAt:    pc.CreatedAt,
	}
}

func ParseToArr(arr []Post) []domain.Post {
	var res []domain.Post

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func ParsePostCommentToArr(arr []PostComent) []domain.PostComent {
	var res []domain.PostComent

	for _, val := range arr {
		res = append(res, val.ToPostComent())
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

func FromModel(data domain.Post) Post {
	var res Post
	res.ID = uint(data.ID)
	res.Userid = data.Userid
	res.Photo = data.Photo
	res.Caption = data.Caption
	res.CreatedAt = data.CreatedAt
	return res
}
