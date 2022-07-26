package usecase

import (
	"backend/domain"

	"github.com/go-playground/validator"
)

type postUsecase struct {
	postData domain.PostData
	validate *validator.Validate
}

func New(pd domain.PostData) domain.PostUseCase {
	return &postUsecase{
		postData: pd,
	}
}

// CreatePost implements domain.PostUseCase
func (puc *postUsecase) CreatePost(newpost domain.Post, userid int) int {
	panic("unimplemented")
}

// UpdatePost implements domain.PostUseCase
func (puc *postUsecase) UpdatePost(newpost domain.Post, userid int) int {
	panic("unimplemented")
}
