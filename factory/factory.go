package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "backend/features/User/data"
	udeli "backend/features/User/delivery"
	uc "backend/features/User/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	userCase := uc.New(userData, validator)
	userHandler := udeli.New(userCase)
	udeli.RouteUser(e, userHandler)
}
