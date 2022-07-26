package data

import (
	"backend/domain"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.PostData {
	return &postData{
		db: db,
	}
}

// CreatePostData implements domain.PostData
func (pd *postData) CreatePostData(newpost domain.Post) domain.User {
	panic("unimplemented")
}

// UpdatePostData implements domain.PostData
func (pd *postData) UpdatePostData(newpost domain.Post) domain.User {
	panic("unimplemented")
}
