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

func (pd *postData) CheckUser(newpost domain.Post) string {
	var post = FromModel(newpost)
	var name string
	err := pd.db.Model(&Post{}).Select("users.username").Joins("join users on posts.userid = users.id").
		Where("userid = ?", post.Userid).Limit(1).Scan(&name)

	if err.Error != nil {
		log.Println("error find data")
		return ""
	}

	return name
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

func (pd *postData) ReadAllCommentData() []domain.CommentUser {
	var com []CommentUser

	err2 := pd.db.Model(&Post{}).Order("comments.id DESC").Select("comments.id, users.firstname, users.lastname, users.photoprofile, comments.postid, comments.comment, comments.created_at").
		Joins("join comments on comments.postid = posts.id ").Joins("join users on comments.userid = users.id ").Find(&com).Limit(50)

	if err2.Error != nil {
		log.Println("Cannot retrieve object", err2.Error)
		return nil
	}

	return ParseCommentUserToArr(com)
}

func (pd *postData) ReadAllPostData() []domain.PostComent {
	var tmp []PostComent
	err := pd.db.Model(&Post{}).Order("posts.id DESC").Select("posts.id, users.firstname, users.lastname, users.username, users.photoprofile, posts.photo, posts.caption, posts.created_at").
		Joins("left join users on users.ID = posts.userid").Find(&tmp).Limit(50).Error

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

func (pd *postData) DeletePostData(postid, userid int) bool {
	res := pd.db.Delete(&Post{}, "ID = ? AND userid = ?", postid, userid)

	if res.Error != nil {
		log.Println("Cannot delete data", res)
		return false
	}
	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}

	return true
}
