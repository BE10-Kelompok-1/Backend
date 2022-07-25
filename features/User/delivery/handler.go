package delivery

import (
	"backend/domain"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	useUsecase domain.UserUseCase
}

func New(uuc domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		useUsecase: uuc,
	}
}

// Register implements domain.UserHandler
func (*userHandler) Register() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements domain.UserHandler
func (*userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
