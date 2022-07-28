package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Post struct {
	ID        int
	Userid    int
	Photo     string
	Caption   string
	CreatedAt time.Time
	Comments  []Comment
}

type PostComent struct {
	ID           int
	Firstname    string
	Lastname     string
	Username     string
	Photoprofile string
	Photo        string
	Caption      string
	CreatedAt    time.Time
}

//line
type PostHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	ReadAll() echo.HandlerFunc
}

type PostUseCase interface {
	CreatePost(newpost Post, userid int) int
	UpdatePost(newpost Post, postid, userid int) int
	ReadAllPost() ([]PostComent, []CommentUser, int)
}

type PostData interface {
	CreatePostData(newpost Post) Post
	UpdatePostData(newpost Post) Post
	ReadAllPostData() []PostComent
	ReadAllCommentData() []CommentUser
}
