package delivery

import (
	"backend/domain"

	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postUseCase domain.PostUseCase
}

func New(puc domain.PostUseCase) domain.PostHandler {
	return &postHandler{
		postUseCase: puc,
	}
}

// Create implements domain.PostHandler
func (*postHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements domain.PostHandler
func (*postHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
