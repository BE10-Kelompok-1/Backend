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

func (puc *postUsecase) ReadAllPost() ([]domain.Post, int) {
	reads := puc.postData.ReadAllPostData()

	if len(reads) == 0 {
		return nil, 404
	}

	return reads, 200
}

func (puc *postUsecase) ReadMyPost(userid int) ([]domain.Post, int) {
	read := puc.postData.ReadMyPostData(userid)

	if len(read) == 0 {
		return nil, 404
	}

	return read, 200
}
