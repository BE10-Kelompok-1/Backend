package delivery

import (
	"backend/domain"
	"backend/features/common"
	"net/http"
	"strings"

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

func (uh *userHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv := c.Param("username")
		data, err := uh.useUsecase.SearchUser(cnv)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"photoprofile": data.Photoprofile,
			"firstname":    data.Firstname,
			"lastname":     data.Lastname,
			"username":     data.Username,
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		data, err := uh.useUsecase.DeleteUser(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user",
		})
	}
}
