package domain

import (
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID      int
	Userid  int
	Postid  int
	Comment string
}

type CommentHandler interface {
	Create() echo.HandlerFunc
}

type CommentUseCase interface {
	CreateComment(newcomment Comment, userid int) int
}

type CommentData interface {
	CreateCommentData(newcomment Comment) Comment
}
