package usecase

import (
	"backend/domain"

	"github.com/go-playground/validator/v10"
)

type commentUsecase struct {
	commentData domain.CommentData
	validate    *validator.Validate
}

func New(pd domain.CommentData, v *validator.Validate) domain.CommentUseCase {
	return &commentUsecase{
		commentData: pd,
		validate:    v,
	}
}

// CreateComment implements domain.CommentUseCase
func (*commentUsecase) CreateComment(newcomment domain.Comment, postid int, userid int) int {
	panic("unimplemented")
}
