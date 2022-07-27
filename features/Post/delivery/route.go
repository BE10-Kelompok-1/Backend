package delivery

import (
	"backend/config"
	"backend/domain"
	"backend/features/Post/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uh domain.PostHandler) {
	e.Pre(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	post := e.Group("/post")
	post.GET("", uh.ReadAll())
	post.GET("/userpost", uh.ReadAll(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
