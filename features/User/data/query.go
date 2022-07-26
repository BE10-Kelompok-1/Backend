package data

import (
	"backend/domain"
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

func (ud *userData) SearchUserData(username string) (domain.User, error) {
	var tmp User
	err := ud.db.Where("username = ?", username).First(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}, err
	}
	return tmp.ToModel(), nil
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
