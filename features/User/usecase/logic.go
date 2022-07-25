package usecase

import (
	"backend/domain"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

// Register implements domain.UserUseCase
func (*userUseCase) Register(newuser domain.User, cost int) (domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUseCase
func (*userUseCase) UpdateUser(newuser domain.User, userid int, cost int) (domain.User, error) {
	panic("unimplemented")
}

func New(uuc domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: uuc,
		validate: v,
	}
}

func (uuc *userUseCase) SearchUser(username string) (domain.User, error) {
	data, err := uuc.userData.SearchUserData(username)

	if err != nil {
		log.Println("Use Case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}
	return data, nil
}

func (uuc *userUseCase) DeleteUser(id int) (bool, error) {
	data := uuc.userData.DeleteUserData(id)

	if !data {
		return false, errors.New("failed delete")
	}

	return true, nil
}
