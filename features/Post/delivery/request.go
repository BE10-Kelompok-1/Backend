package delivery

import (
	"backend/domain"
)

type PostFormat struct {
	Userid  int
	Photo   string `json:"photo" form:"photo"`
	Caption string `json:"caption" form:"caption" validate:"required"`
}

func (i *PostFormat) ToModel() domain.Post {
	return domain.Post{
		Userid:  i.Userid,
		Photo:   i.Photo,
		Caption: i.Caption,
	}
}
