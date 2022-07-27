package delivery

import (
	"backend/domain"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentUseCase domain.CommentUseCase
}

func New(cuc domain.CommentUseCase) domain.CommentHandler {
	return &commentHandler{
		commentUseCase: cuc,
	}
}

// Create implements domain.CommentHandler
func (ch *commentHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}
