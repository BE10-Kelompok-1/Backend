package data

import (
	"backend/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Userid  int
	Postid  int
	Comment string `json:"comment" form:"comment" validate:"required"`
}

func (c *Comment) ToModel() domain.Comment {
	return domain.Comment{
		ID:      int(c.ID),
		Userid:  c.Userid,
		Postid:  c.Postid,
		Comment: c.Comment,
	}
}

func ParseToArr(arr []Comment) []domain.Comment {
	var res []domain.Comment

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.Comment) Comment {
	var res Comment
	res.ID = uint(data.ID)
	res.Userid = data.Userid
	res.Postid = data.Postid
	res.Comment = data.Comment
	return res
}
