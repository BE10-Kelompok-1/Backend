package usecase

import (
	"backend/domain"
	"backend/features/User/data"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(uuc domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: uuc,
		validate: v,
	}
}

func (uuc *userUseCase) SearchUser(username string) (domain.User, []domain.UserPosting, []domain.CommentUser, int) {
	profile := uuc.userData.SearchUserData(username)
	posting := uuc.userData.SearchUserPostingData(username)
	comment := uuc.userData.SearchUserPostingCommentData(username)

	if username == "" {
		log.Println("Wrong input")
		return domain.User{}, nil, nil, 404
	}

	if profile.ID == 0 {
		log.Println("Data not found")
		return domain.User{}, nil, nil, 404
	}
	return profile, posting, comment, 200
}

func (uuc *userUseCase) DeleteUser(id int) int {
	delete := uuc.userData.DeleteUserData(id)

	if !delete {
		log.Println("Data not found")
		return 404
	}

	return 200
}

// Register implements domain.UserUseCase
func (uuc *userUseCase) RegisterUser(newuser domain.User, cost int) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 400
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return 500
	}
	user.Password = string(hashed)
	insert := uuc.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return 404
	}

	return 200
}

// UpdateUser implements domain.UserUseCase
func (uuc *userUseCase) UpdateUser(newuser domain.User, userid int, cost int) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)

	if userid == 0 {
		log.Println("Data not found")
		return 404
	}

	if validError != nil {
		log.Println("Validation errror : ", validError.Error())
		return 400
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 400
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Println("Error encrypt password", err)
		return 500
	}

	user.ID = uint(userid)
	user.Password = string(hashed)

	uuc.userData.UpdateUserData(user.ToModel())

	return 200
}

func (uuc *userUseCase) LoginUser(userdata domain.User) (domain.User, error) {
	login := uuc.userData.LoginData(userdata)

	if login.ID == 0 {
		return domain.User{}, errors.New("no data")
	}

	hashpw := uuc.userData.GetPasswordData(userdata.Username)

	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userdata.Password))

	if err != nil {
		log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
		return domain.User{}, err
	}

	return login, nil
}

func (uuc *userUseCase) ProfileUser(userid int) (domain.User, []domain.UserPosting, []domain.CommentUser, int) {
	profile := uuc.userData.ProfileUserData(userid)
	posting := uuc.userData.GetUserPostingData(userid)
	comment := uuc.userData.GetUserCommentData(userid)

	if userid == 0 {
		log.Println("Need login or register first")
		return domain.User{}, nil, nil, 404
	}

	if profile.ID == 0 {
		log.Println("Data not found")
		return domain.User{}, nil, nil, 404
	}
	return profile, posting, comment, 200
}
