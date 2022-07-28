package usecase

import (
	"backend/domain"
	"backend/domain/mocks"
	"testing"
)

func TestCreatePost(t *testing.T) {
	repo := new(mocks.PostData)

	mockData := domain.Post{Photo: ""}
}
