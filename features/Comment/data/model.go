package data

import (
	"backend/domain"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Userid  int
	Postid  int
	Comment string `json:"comment" form:"comment" validate:"required"`
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

func (c *Comment) ToModel() domain.Comment {
	return domain.Comment{
		ID:      int(c.ID),
		Userid:  c.Userid,
		Postid:  c.Postid,
		Comment: c.Comment,
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

func ParseToArr(arr []Comment) []domain.Comment {
	var res []domain.Comment

	for _, val := range arr {
		res = append(res, val.ToModel())
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

func FromModel(data domain.Comment) Comment {
	var res Comment
	res.ID = uint(data.ID)
	res.Userid = data.Userid
	res.Postid = data.Postid
	res.Comment = data.Comment
	return res
}
