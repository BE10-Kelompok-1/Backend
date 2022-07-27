package delivery

import (
	"backend/domain"
	"backend/features/common"
	"log"
	"net/http"

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
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		bind := c.Bind(&newuser)
		cost := 10

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := uh.useUsecase.RegisterUser(newuser.ToModel(), cost)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Register success",
		})
	}
}

// Update implements domain.UserHandler

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		cost := 10
		id := common.ExtractData(c)
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := uh.useUsecase.UpdateUser(newuser.ToModel(), id, cost)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}

func (uh *userHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv := c.Param("username")
		data, status := uh.useUsecase.SearchUser(cnv)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"photoprofile": data.Photoprofile,
			"firstname":    data.Firstname,
			"lastname":     data.Lastname,
			"username":     data.Username,
			"code":         status,
			"message":      "get data success",
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := common.ExtractData(c)
		status := uh.useUsecase.DeleteUser(id)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Success delete data",
		})
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var datalogin LoginFormat
		bind := c.Bind(&datalogin)

		if bind != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		data, err := uh.useUsecase.LoginUser(datalogin.ToModelLogin())

		if err != nil {
			log.Println("Login failed", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong username or password",
			})
		}

		token := common.GenerateToken(int(data.ID))

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Login success",
			"token":   token,
		})
	}
}
