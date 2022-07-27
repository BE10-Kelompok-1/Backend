package data

import (
	"backend/domain"

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
	panic("unimplemented")
}
