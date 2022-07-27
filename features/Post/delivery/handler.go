package delivery

import (
	"backend/domain"
	"backend/features/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postUseCase domain.PostUseCase
}

func New(puc domain.PostUseCase) domain.PostHandler {
	return &postHandler{
		postUseCase: puc,
	}
}

// Create implements domain.PostHandler
func (*postHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements domain.PostHandler
func (*postHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (ph *postHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, status := ph.postUseCase.ReadAllPost()

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
			"data":    data,
			"code":    status,
			"message": "get data success",
		})
	}
}

func (ph *postHandler) ReadMy() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		data, status := ph.postUseCase.ReadMyPost(id)

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
			"data": data,
		})

	}
}
