package data

import (
	"backend/domain"
	"log"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CommentData {
	return &commentData{
		db: db,
	}
}

// CreateCommentData implements domain.CommentData
func (cd *commentData) CreateCommentData(newcomment domain.Comment) domain.Comment {
	var comment = FromModel(newcomment)
	err := cd.db.Create(&comment)

	if err.Error != nil {
		log.Println("Cant create user object", err.Error)
		return domain.Comment{}
	}

	return comment.ToModel()
}
