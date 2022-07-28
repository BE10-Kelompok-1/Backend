package usecase

import (
	"backend/domain"
	"backend/features/Comment/data"
	"log"

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
func (cuc *commentUsecase) CreateComment(newcomment domain.Comment, userid int) int {
	var comment = data.FromModel(newcomment)
	validError := cuc.validate.Struct(comment)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	comment.Userid = userid
	create := cuc.commentData.CreateCommentData(comment.ToModel())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}

	return 200
}

func (cuc *commentUsecase) ReadComment() ([]domain.CommentUser, int) {
	reads := cuc.commentData.ReadCommentData()

	if len(reads) == 0 {
		return nil, 404
	}

	return reads, 200
}
