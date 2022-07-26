package delivery

import (
	"backend/domain"
	"backend/features/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentUseCase domain.CommentUseCase
}

func New(cuc domain.CommentUseCase) domain.CommentHandler {
	return &commentHandler{
		commentUseCase: cuc,
	}
}

// Create implements domain.CommentHandler
func (ch *commentHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newcomment CommentFormat
		id := common.ExtractData(c)
		bind := c.Bind(&newcomment)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ch.commentUseCase.CreateComment(newcomment.ToModel(), id)

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

func (ch *commentHandler) Read() echo.HandlerFunc {
	return func(c echo.Context) error {
		comment, status := ch.commentUseCase.ReadComment()

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
			"comments": comment,
			"code":     status,
			"message":  "get data success",
		})
	}
}

func (ch *commentHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		param := c.Param("commentid")
		id := common.ExtractData(c)
		cnv, err := strconv.Atoi(param)

		if err != nil {
			log.Println("cant convert to int", err)
			return c.JSON(http.StatusInternalServerError, "cant convert to int")
		}

		status := ch.commentUseCase.DeleteComment(cnv, id)

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
			"message": "Success delete comment",
		})
	}
}
