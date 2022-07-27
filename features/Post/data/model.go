package data

import (
	"backend/domain"
	"backend/features/Comment/data"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Userid   int
	Photo    string         `json:"photo" form:"photo"`
	Caption  string         `json:"caption" form:"caption" validate:"required"`
	Comments []data.Comment `gorm:"foreignKey:Postid"`
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

func ParseToArr(arr []Post) []domain.Post {
	var res []domain.Post

	for _, val := range arr {
		res = append(res, val.ToModel())
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
