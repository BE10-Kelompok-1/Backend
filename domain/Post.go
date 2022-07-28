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
	Comments     []CommentUser
	// Idcoment           int
	// Firstnamecoment    string
	// Lastnamecoment     string
	// Photoprofilecoment string
	// Comment            string
	// Created_at         time.Time
}

type PostHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	ReadAll() echo.HandlerFunc
	ReadMy() echo.HandlerFunc
}

type PostUseCase interface {
	CreatePost(newpost Post, userid int) int
	UpdatePost(newpost Post, postid, userid int) int
	ReadAllPost() ([]PostComent, []CommentUser, int)
	ReadMyPost(userid int) ([]Post, int)
}

type PostData interface {
	CreatePostData(newpost Post) Post
	UpdatePostData(newpost Post) Post
	ReadAllPostData() []PostComent
	ReadMyPostData(userid int) []Post
	ReadAllCommentData() []CommentUser
}
