package usecase

import (
	"backend/domain"
	"backend/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePost(t *testing.T) {
	repo := new(mocks.PostData)
	mockData := domain.Post{ID: 1, Userid: 2, Photo: "jpg", Caption: "siap"}
	emptyMockData := domain.Post{ID: 0, Userid: 0, Photo: "", Caption: ""}
	t.Run("Success insert post", func(t *testing.T) {
		repo.On("CreatePostData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreatePost(mockData, 1)

		assert.Equal(t, 200, res)
		assert.Greater(t, mockData.ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CreatePostData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreatePost(emptyMockData, 5)

		assert.Equal(t, 400, res)
		assert.Equal(t, emptyMockData.ID, 0)
	})
}

func TestDeletePost(t *testing.T) {
	repo := new(mocks.PostData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeletePostData", mock.Anything, mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeletePost(0, 1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeletePostData", mock.Anything, mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeletePost(1, 1)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}

func TestReadAllPost(t *testing.T) {
	repo := new(mocks.PostData)
	returnData1 := []domain.PostComent{{ID: 1, Firstname: "Vanili", Lastname: "Nugroho", Username: "vanili", Photoprofile: "jpg", Photo: "jpg", Caption: "a"}}
	returnData2 := []domain.CommentUser{{Id: 1, Firstname: "Vanili", Lastname: "Nugroho", Photoprofile: "apa", Postid: 2, Comment: "apa tuh"}}
	t.Run("Success get all Post", func(t *testing.T) {
		repo.On("ReadAllCommentData").Return(returnData2).Once()
		repo.On("ReadAllPostData").Return(returnData1).Once()
		usecase := New(repo, validator.New())
		res, res2, status := usecase.ReadAllPost()
		assert.Equal(t, 200, status)
		assert.GreaterOrEqual(t, len(res), 1)
		assert.GreaterOrEqual(t, len(res2), 1)
		assert.Greater(t, res[0].ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("No data found", func(t *testing.T) {
		repo.On("ReadAllCommentData").Return([]domain.CommentUser{}).Once()
		repo.On("ReadAllPostData").Return([]domain.PostComent{}).Once()
		usecase := New(repo, validator.New())
		res, res2, status := usecase.ReadAllPost()
		assert.Equal(t, 404, status)
		assert.Equal(t, len(res), 0)
		assert.Equal(t, len(res2), 0)
		assert.Equal(t, []domain.PostComent([]domain.PostComent(nil)), res)
		assert.Equal(t, []domain.PostComent(nil), res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.PostData)
	mockData := domain.Post{Photo: "jpg", Caption: "siap"}
	mockData2 := domain.Post{ID: 1, Userid: 2, Photo: "jpg", Caption: "siap"}
	emptyMockData := domain.Post{ID: 0, Userid: 0, Photo: "", Caption: ""}

	t.Run("Success Update Post", func(t *testing.T) {
		repo.On("UpdatePostData", mock.Anything).Return(mockData2).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdatePost(mockData, 1, 1)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("UpdatePostData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdatePost(mockData, 0, 0)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
}
