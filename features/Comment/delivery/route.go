package delivery

import (
	"backend/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteComment(e *echo.Echo, ph domain.PostHandler) {
	e.Pre(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

}
