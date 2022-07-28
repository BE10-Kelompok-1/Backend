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

func (cd *commentData) ReadCommentData() []domain.CommentUser {
	var data []CommentUser
	err := cd.db.Model(&Comment{}).Order("comments.id DESC").Select("comments.id, users.firstname, users.lastname, users.photoprofile, comments.postid, comments.comment, comments.created_at").
		Joins("left join users on users.ID = comments.userid").Find(&data).Limit(50)
	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}
	return ParseCommentUserToArr(data)
}
