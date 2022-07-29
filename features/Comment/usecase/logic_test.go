package usecase

import (
	"backend/domain"
	"backend/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteComment(t *testing.T) {
	repo := new(mocks.CommentData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteCommentData", mock.Anything, mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteComment(0, 1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeleteCommentData", mock.Anything, mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteComment(1, 1)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}

func TestReadComment(t *testing.T) {
	repo := new(mocks.CommentData)
	returnData := []domain.CommentUser{{Id: 1, Firstname: "Vanili", Lastname: "Nugroho", Photoprofile: "apa", Postid: 2, Comment: "apa tuh"}}

	t.Run("Success get all comment", func(t *testing.T) {
		repo.On("ReadCommentData").Return(returnData).Once()
		usecase := New(repo, validator.New())
		res, status := usecase.ReadComment()
		assert.Equal(t, 200, status)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.Greater(t, res[0].Id, 0)
		repo.AssertExpectations(t)
	})

	t.Run("No data found", func(t *testing.T) {
		repo.On("ReadCommentData").Return([]domain.CommentUser{}).Once()
		usecase := New(repo, validator.New())
		res, status := usecase.ReadComment()
		assert.Equal(t, 404, status)
		assert.Equal(t, len(res), 0)
		assert.Equal(t, []domain.CommentUser(nil), res)
		repo.AssertExpectations(t)
	})
}

func TestCreateComment(t *testing.T) {
	repo := new(mocks.CommentData)
	mockData := domain.Comment{ID: 1, Userid: 2, Postid: 3, Comment: "males"}
	emptyMockData := domain.Comment{ID: 0, Userid: 0, Postid: 0, Comment: ""}
	t.Run("Success insert comment", func(t *testing.T) {
		repo.On("CreateCommentData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateComment(mockData, 1)

		assert.Equal(t, 200, res)
		assert.Greater(t, mockData.ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CreateCommentData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateComment(emptyMockData, 5)

		assert.Equal(t, 400, res)
		assert.Equal(t, emptyMockData.ID, 0)
	})
}
