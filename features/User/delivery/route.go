package delivery

import (
	"backend/config"
	"backend/domain"
	"backend/features/User/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, bc domain.UserHandler) {
	e.GET("/user/:username", bc.Search())
	e.DELETE("/user", bc.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
