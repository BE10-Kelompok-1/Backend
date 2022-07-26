package data

import (
	"backend/domain"
	"log"

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
func (pd *postData) CreatePostData(newpost domain.Post) domain.Post {
	var post = FromModel(newpost)
	err := pd.db.Create(&post)

	if err.Error != nil {
		log.Println("Cant create user object", err.Error)
		return domain.Post{}
	}

	return post.ToModel()
}

// UpdatePostData implements domain.PostData
func (pd *postData) UpdatePostData(newpost domain.Post) domain.Post {
	var post = FromModel(newpost)
	err := pd.db.Model(&Post{}).Where("ID = ?", post.ID).Updates(post)

	if err.Error != nil {
		log.Println("Cant update post object", err.Error.Error())
		return domain.Post{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.Post{}
	}

	return post.ToModel()
}
