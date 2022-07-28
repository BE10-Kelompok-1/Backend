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
	err := pd.db.Model(&Post{}).Where("ID = ? AND Userid = ?", post.ID, post.Userid).Updates(post)

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

func (pd *postData) ReadAllPostData() []domain.PostComent {
	var tmp []PostComent
	err := pd.db.Find(&tmp).Error
	//err := pd.db.Model(&Post{}).Select("comments.id, users.firstname, users.lastname, users.username, users.photoprofile, posts.photo, posts.caption, posts.created_at, comments.id, comments.comment, comments.created_at").Joins("left join users on users.ID = comments.userid").Find(&data)

	if err != nil {
		log.Println("Cannot retrieve object", err.Error())
		return nil
	}

	if len(tmp) == 0 {
		log.Println("No data found", gorm.ErrRecordNotFound.Error())
		return nil
	}

	return ParsePostCommentToArr(tmp)
}

func (pd *postData) ReadMyPostData(userid int) []domain.Post {
	var tmp []Post
	err := pd.db.Where("Userid = ?", userid).Find(&tmp).Error

	if err != nil {
		log.Println("There is problem with data")
		return nil
	}
	return ParseToArr(tmp)
}
