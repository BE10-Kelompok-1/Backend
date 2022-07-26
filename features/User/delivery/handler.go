package delivery

import (
	"backend/domain"
	"backend/features/common"
	awss3 "backend/infrastructure/database/aws"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	useUsecase domain.UserUseCase
	userdata   domain.UserData
	conn       *session.Session
}

func New(uuc domain.UserUseCase, ud domain.UserData, aws *session.Session) domain.UserHandler {
	return &userHandler{
		useUsecase: uuc,
		userdata:   ud,
		conn:       aws,
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

		file, err := c.FormFile("photoprofile")

		if err != nil {
			log.Println(err)
		}

		filename := fmt.Sprintf("%s_profilepic.jpg", newuser.Username)
		log.Println(filename)
		link := awss3.DoUpload(uh.conn, *file, filename)
		newuser.Photoprofile = link

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

		file, err := c.FormFile("photoprofile")

		if err == nil {
			log.Println(err)
			filename := fmt.Sprintf("%s_profilepic.jpg", newuser.Username)
			log.Println(filename)
			link := awss3.DoUpload(uh.conn, *file, filename)
			newuser.Photoprofile = link
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
		profile, posting, comment, status := uh.useUsecase.SearchUser(cnv)

		var comarrmap = []domain.CommentUser{}
		var arrmap []map[string]interface{}
		var arrmap2 []map[string]interface{}

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
		var res2 = map[string]interface{}{}
		res2["photoprofile"] = profile.Photoprofile
		res2["firstname"] = profile.Firstname
		res2["lastname"] = profile.Lastname
		res2["username"] = profile.Username
		for i := 0; i < len(posting); i++ {
			var res = map[string]interface{}{}
			for j := 0; j < len(comment); j++ {
				if posting[i].ID == comment[j].Postid {
					comarrmap = append(comarrmap, comment[j])
				}
			}
			res["id"] = posting[i].ID
			res["photo"] = posting[i].Photo
			res["caption"] = posting[i].Caption
			res["created_at"] = posting[i].CreatedAt
			res["comments"] = comarrmap

			comarrmap = comarrmap[len(comarrmap):]
			arrmap = append(arrmap, res)
			log.Println("arrmap postkomen", arrmap)

		}
		res2["posts"] = arrmap
		arrmap2 = append(arrmap2, res2)

		return c.JSON(http.StatusOK, arrmap2)
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

func (uh *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		profile, posting, comment, status := uh.useUsecase.ProfileUser(id)

		var comarrmap = []domain.CommentUser{}
		var arrmap []map[string]interface{}
		var arrmap2 []map[string]interface{}

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
		var res2 = map[string]interface{}{}
		res2["photoprofile"] = profile.Photoprofile
		res2["firstname"] = profile.Firstname
		res2["lastname"] = profile.Lastname
		res2["username"] = profile.Username
		res2["email"] = profile.Email
		res2["birthdate"] = profile.Birthdate
		res2["password"] = profile.Password
		for i := 0; i < len(posting); i++ {
			var res = map[string]interface{}{}
			for j := 0; j < len(comment); j++ {
				if posting[i].ID == comment[j].Postid {
					comarrmap = append(comarrmap, comment[j])
				}
			}
			res["id"] = posting[i].ID
			res["photo"] = posting[i].Photo
			res["caption"] = posting[i].Caption
			res["created_at"] = posting[i].CreatedAt
			res["comments"] = comarrmap

			comarrmap = comarrmap[len(comarrmap):]
			arrmap = append(arrmap, res)
			log.Println("arrmap postkomen", arrmap)

		}
		res2["posts"] = arrmap
		arrmap2 = append(arrmap2, res2)

		return c.JSON(http.StatusOK, arrmap2)
	}

}
