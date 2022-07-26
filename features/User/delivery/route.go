package delivery

import (
	"backend/config"
	"backend/domain"
	"backend/features/User/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uh domain.UserHandler) {
	// e.Pre(middleware.CORS())
	// e.Use(middleware.RemoveTrailingSlash())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	user := e.Group("/users")
	user.GET("/:username", uh.Search())
	user.DELETE("", uh.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	user.POST("", uh.Register())
	user.PUT("", uh.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
