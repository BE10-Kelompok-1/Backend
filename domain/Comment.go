package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID      int
	Userid  int
	Postid  int
	Comment string
}

type CommentUser struct {
	Id           int
	Firstname    string
	Lastname     string
	Photoprofile string
	Postid       int
	Comment      string
	Created_at   time.Time
}

type CommentHandler interface {
	Create() echo.HandlerFunc
	Read() echo.HandlerFunc
}

type CommentUseCase interface {
	CreateComment(newcomment Comment, userid int) int
	ReadComment() ([]CommentUser, int)
}

type CommentData interface {
	CreateCommentData(newcomment Comment) Comment
	ReadCommentData() []CommentUser
}
