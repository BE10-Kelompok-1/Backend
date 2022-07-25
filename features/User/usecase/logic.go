package usecase

import (
	"backend/domain"

	"github.com/go-playground/validator"
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
