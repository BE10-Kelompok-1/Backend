package usecase

import (
	"backend/domain"
	"backend/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	emptyMockData := domain.User{ID: 0, Firstname: "", Lastname: "", Username: "",
		Email: "", Password: "", Birthdate: "", Photoprofile: ""}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru"

	invalidData := mockData
	invalidData.Email = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success insert data", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validator Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(noData, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	returnData := domain.User{ID: 1, Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	mockData := domain.User{Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	invalidData := mockData
	invalidData.Firstname = ""

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validator Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestSearchUser(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := domain.User{ID: 1, Firstname: "Vanilia", Lastname: "Nugroho", Username: "vaniliacahya",
		Email: "vanili@vanili", Password: "vanili123", Birthdate: "2022-07-20", Photoprofile: "vanili.jpg"}

	returnData2 := domain.User{ID: 0, Firstname: "", Lastname: "", Username: "",
		Email: "", Password: "", Birthdate: "", Photoprofile: ""}

	t.Run("Succes get user", func(t *testing.T) {
		repo.On("SearchUserData", mock.Anything).Return(returnData).Once()
		usecase := New(repo, validator.New())
		search, res := usecase.SearchUser("vaniliacahya")

		assert.Equal(t, 200, res)
		assert.Greater(t, search.ID, 0)
		assert.Equal(t, "Vanilia", search.Firstname)
		assert.Equal(t, "Nugroho", search.Lastname)
		assert.Equal(t, "vaniliacahya", search.Username)
		assert.Equal(t, "vanili@vanili", search.Email)
		assert.Equal(t, "vanili123", search.Password)
		assert.Equal(t, "2022-07-20", search.Birthdate)
		assert.Equal(t, "vanili.jpg", search.Photoprofile)
		repo.AssertExpectations(t)
	})

	t.Run("No data found", func(t *testing.T) {
		repo.On("SearchUserData", mock.Anything).Return(returnData2)
		usecase := New(repo, validator.New())
		search, res := usecase.SearchUser("vanilii")
		assert.Equal(t, 404, res)
		assert.Equal(t, 0, search.ID)
		assert.Equal(t, "", search.Firstname)
		assert.Equal(t, "", search.Lastname)
		assert.Equal(t, "", search.Username)
		assert.Equal(t, "", search.Email)
		assert.Equal(t, "", search.Password)
		assert.Equal(t, "", search.Birthdate)
		assert.Equal(t, "", search.Photoprofile)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteUserData", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteUser(1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeleteUserData", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteUser(100)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}
