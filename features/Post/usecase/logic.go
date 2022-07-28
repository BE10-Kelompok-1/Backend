package usecase

import (
	"backend/domain"
	"backend/features/Post/data"
	"log"

	"github.com/go-playground/validator/v10"
)

type postUsecase struct {
	postData domain.PostData
	validate *validator.Validate
}

func New(pd domain.PostData, v *validator.Validate) domain.PostUseCase {
	return &postUsecase{
		postData: pd,
		validate: v,
	}
}

// CreatePost implements domain.PostUseCase
func (puc *postUsecase) CreatePost(newpost domain.Post, userid int) int {
	var post = data.FromModel(newpost)
	validError := puc.validate.Struct(post)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	post.Userid = userid
	create := puc.postData.CreatePostData(post.ToModel())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}

	return 200
}

// UpdatePost implements domain.PostUseCase
func (puc *postUsecase) UpdatePost(newpost domain.Post, postid, userid int) int {
	var post = data.FromModel(newpost)
	validError := puc.validate.Struct(post)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}
	post.ID = uint(postid)
	post.Userid = userid
	update := puc.postData.UpdatePostData(post.ToModel())

	if update.ID == 0 {
		log.Println("Empty Data")
		return 404
	}

	return 200
}

func (puc *postUsecase) ReadAllPost() ([]domain.PostComent, int) {
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
