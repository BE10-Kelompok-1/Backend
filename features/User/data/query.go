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
	panic("unimplemented")
}

// UpdateUserData implements domain.UserData
func (ud *userData) UpdateUserData(newuser domain.User) domain.User {
	panic("unimplemented")
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
