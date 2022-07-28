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
	Photo    string             `json:"photo" form:"photo"`
	Caption  string             `json:"caption" form:"caption" validate:"required"`
	Comments []data.CommentUser `gorm:"foreignKey:Postid"`
}

type PostComent struct {
	ID           int
	Firstname    string
	Lastname     string
	Photoprofile string
	Comments     []data.CommentUser `gorm:"foreignKey:Postid"`
	CreatedAt    time.Time
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

func (pc *PostComent) ToPostComent() domain.PostComent {
	return domain.PostComent{
		ID:           pc.ID,
		Firstname:    pc.Firstname,
		Lastname:     pc.Lastname,
		Photoprofile: pc.Photoprofile,
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

func FromModel(data domain.Post) Post {
	var res Post
	res.ID = uint(data.ID)
	res.Userid = data.Userid
	res.Photo = data.Photo
	res.Caption = data.Caption
	res.CreatedAt = data.CreatedAt
	return res
}
