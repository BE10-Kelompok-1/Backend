package delivery

import (
	"backend/config"
	"backend/domain"
	"backend/features/Post/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePost(e *echo.Echo, ph domain.PostHandler) {
	e.Pre(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	post := e.Group("/posts")
	post.POST("", ph.Create(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	post.PUT("", ph.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	post.GET("", ph.ReadAll())
	post.GET("/userpost", ph.ReadAll(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
