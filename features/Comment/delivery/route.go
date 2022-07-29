package delivery

import (
	"backend/config"
	"backend/domain"
	"backend/features/Comment/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteComment(e *echo.Echo, ch domain.CommentHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	comment := e.Group("/comments")
	comment.POST("", ch.Create(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	comment.GET("", ch.Read())
	comment.DELETE("/:commentid", ch.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
