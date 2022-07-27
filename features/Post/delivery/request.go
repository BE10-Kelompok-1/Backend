package delivery

import "backend/domain"

type PostFormat struct {
	Userid  int
	Photo   string `json:"Photo"`
	Caption string `json:"Caption" validate:"required"`
}

func (i *PostFormat) ToModel() domain.Post {
	return domain.Post{
		Userid:  i.Userid,
		Photo:   i.Photo,
		Caption: i.Caption,
	}
}
