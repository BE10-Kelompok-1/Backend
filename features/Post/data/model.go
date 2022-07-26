package data

import (
	"backend/domain"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User    domain.User `gorm:"foreignKey:id;references:UserID"`
	Photo   string      `json:"photo"`
	Caption string      `json:"caption" validate:"required"`
}

func (p *Post) ToModel() domain.Post {
	return domain.Post{
		ID:        int(p.ID),
		User:      p.User,
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
	res.User = data.User
	res.Photo = data.Photo
	res.Caption = data.Caption
	res.CreatedAt = data.CreatedAt
	return res
}
