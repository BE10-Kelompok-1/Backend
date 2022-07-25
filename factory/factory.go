package factory

import (
	ud "backend/features/User/data"
	userDelivery "backend/features/User/delivery"
	us "backend/features/User/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)
}
