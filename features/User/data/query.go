package data

import (
	"backend/domain"
	"backend/features/Post/data"
	"log"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) GetUserPostingData(userid int) []domain.UserPosting {
	var tmp []UserPosting
	err := ud.db.Model(&User{}).Order("posts.id DESC").Select("users.ID, posts.ID, posts.photo, posts.caption, posts.created_at").
		Joins("left join posts on posts.userid = users.id").Where("users.iD = ?", userid).Find(&tmp).Limit(50).Error

	if err != nil {
		log.Println("There is problem with data", err.Error())
		return nil
	}

	return ParseUserPostingToArr(tmp)
}

func (ud *userData) GetUserCommentData(userid int) []domain.CommentUser {
	var profile = User{}
	err2 := ud.db.Where("id = ?", userid).First(&profile).Error

	if err2 != nil {
		log.Println("There is problem with data", err2.Error())
		return nil
	}

	var tmp []CommentUser
	err := ud.db.Model(&data.Post{}).Order("comments.id DESC").Select("comments.id, users.firstname, users.lastname, users.photoprofile, comments.postid, comments.comment, comments.created_at").
		Joins("join comments on comments.postid = posts.id").Joins("join users on comments.userid = users.id").Where("posts.userid = ?", profile.ID).Find(&tmp).Limit(50).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return nil
	}
	return ParseCommentUserToArr(tmp)
}

// CheckDuplicate implements domain.UserData
func (ud *userData) CheckDuplicate(newuser domain.User) bool {
	var user User
	err := ud.db.Find(&user, "username = ? OR email = ?", newuser.Username, newuser.Email)

	if err.RowsAffected == 1 {
		log.Println("Duplicated data", err.Error)
		return true
	}

	return false
}

// RegisterData implements domain.UserData
func (ud *userData) RegisterData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	err := ud.db.Create(&user).Error

	if user.ID == 0 {
		log.Println("Invalid ID")
		return domain.User{}
	}

	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

// UpdateUserData implements domain.UserData
func (ud *userData) UpdateUserData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	err := ud.db.Model(&User{}).Where("ID = ?", user.ID).Updates(user)

	if err.Error != nil {
		log.Println("Cant update user object", err.Error.Error())
		return domain.User{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) SearchUserData(username string) domain.User {
	var tmp User
	err := ud.db.Where("username = ?", username).First(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}
	}
	return tmp.ToModel()
}

func (ud *userData) DeleteUserData(userid int) bool {
	res := ud.db.Where("ID = ?", userid).Delete(&User{})
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return false
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}
	return true
}

func (ud *userData) GetPasswordData(username string) string {
	var user User
	err := ud.db.Find(&user, "username = ?", username).Error

	if err != nil {
		log.Println("Cant retrieve user data", err.Error())
		return ""
	}

	return user.Password
}

func (ud *userData) LoginData(userdata domain.User) domain.User {
	var user = FromModel(userdata)
	err := ud.db.First(&user, "username  = ?", userdata.Username).Error

	if err != nil {
		log.Println("Cant login data", err.Error())
		return domain.User{}
	}

	return user.ToModel()

}

func (ud *userData) ProfileUserData(userid int) domain.User {
	var tmp User
	err := ud.db.Where("id = ?", userid).First(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}
	}
	return tmp.ToModel()
}

func (ud *userData) SearchUserPostingData(username string) []domain.UserPosting {
	var tmp []UserPosting
	err := ud.db.Model(&User{}).Order("posts.id DESC").Select("users.ID, posts.ID, posts.photo, posts.caption, posts.created_at").
		Joins("left join posts on posts.userid = users.id").Where("users.username = ?", username).Find(&tmp).Limit(50).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return nil
	}
	return ParseUserPostingToArr(tmp)
}

func (ud *userData) SearchUserPostingCommentData(username string) []domain.CommentUser {
	var profile = User{}
	err2 := ud.db.Where("username = ?", username).First(&profile).Error

	if err2 != nil {
		log.Println("There is problem with data", err2.Error())
		return nil
	}

	var tmp []CommentUser
	err := ud.db.Model(&data.Post{}).Order("comments.id DESC").Select("comments.id, users.firstname, users.lastname, users.photoprofile, comments.postid, comments.comment, comments.created_at").
		Joins("join comments on comments.postid = posts.id").Joins("join users on comments.userid = users.id").Where("posts.userid = ?", profile.ID).Find(&tmp).Limit(50).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return nil
	}
	return ParseCommentUserToArr(tmp)
}
