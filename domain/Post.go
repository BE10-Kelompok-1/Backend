package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Post struct {
	ID        int
	Photo     string
	Caption   string
	CreatedAt time.Time
}

type PostHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type PostUseCase interface {
	CreatePost(newpost Post, userid int) int
	UpdatePost(newpost Post, userid int) int
}

type PostData interface {
	CreatePostData(newpost Post) User
	UpdatePostData(newpost Post) User
}
