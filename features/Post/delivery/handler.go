package delivery

import (
	"backend/domain"
	"backend/features/common"
	awss3 "backend/infrastructure/database/aws"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postData    domain.PostData
	postUseCase domain.PostUseCase
	conn        *session.Session
}

func New(pd domain.PostData, puc domain.PostUseCase, aws *session.Session) domain.PostHandler {
	return &postHandler{
		postUseCase: puc,
		postData:    pd,
		conn:        aws,
	}
}

// Create implements domain.PostHandler
func (ph *postHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newpost PostFormat
		id := common.ExtractData(c)
		bind := c.Bind(&newpost)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		name := ph.postData.CheckUser(newpost.ToModel())

		file, err := c.FormFile("photo")

		if err != nil {
			log.Println(err)
		}

		filename := fmt.Sprintf("%s_postpic", name)
		link := awss3.DoUpload(ph.conn, *file, filename)
		newpost.Photo = link
		status := ph.postUseCase.CreatePost(newpost.ToModel(), id)

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

// Update implements domain.PostHandler
func (ph *postHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newpost PostFormat
		postid := c.Param("postid")
		cnv, err := strconv.Atoi(postid)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong input",
			})
		}

		userid := common.ExtractData(c)
		bind := c.Bind(&newpost)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ph.postUseCase.UpdatePost(newpost.ToModel(), cnv, userid)

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

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}

func (ph *postHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, datacom, status := ph.postUseCase.ReadAllPost()

		var comarrmap = []domain.CommentUser{}
		var arrmap []map[string]interface{}

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

		for i := 0; i < len(data); i++ {
			var res = map[string]interface{}{}
			for j := 0; j < len(datacom); j++ {
				if data[i].ID == datacom[j].Postid {
					comarrmap = append(comarrmap, datacom[j])
				}
			}
			res["id"] = data[i].ID
			res["firstname"] = data[i].Firstname
			res["lastname"] = data[i].Lastname
			res["photoprofile"] = data[i].Photoprofile
			res["photo"] = data[i].Photo
			res["caption"] = data[i].Caption
			res["created_at"] = data[i].CreatedAt
			res["comments"] = comarrmap
			// res["code"] = status
			// res["message"] = "get data success"

			comarrmap = comarrmap[2:]
			arrmap = append(arrmap, res)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"posts":   arrmap,
			"code":    status,
			"message": "get data success",
		})
	}
}
