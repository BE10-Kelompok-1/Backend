package data

import (
	"backend/domain"
	"fmt"
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

func (ud *userData) SearchUserData(username string) (domain.User, []domain.User_Posting) {
	var tmp User
	err := ud.db.Where("username = ?", username).First(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}, []domain.User_Posting{}
	}

	var tmp2 []User_Posting
	err2 := ud.db.Model(&User{}).Select("users.username, posts.id, posts.photo, posts.caption, posts.created_at").
		Joins("left join posts on posts.userid = users.ID").Where("users.username = ?", username).Scan(&tmp2).Error
	if err2 != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}, []domain.User_Posting{}
	}
	fmt.Println(tmp2)
	return tmp.ToModel(), ParseToArr2(tmp2)
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
