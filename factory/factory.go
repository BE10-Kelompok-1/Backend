package factory

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	pd "backend/features/Post/data"
	pdeli "backend/features/Post/delivery"
	pc "backend/features/Post/usecase"

	ud "backend/features/User/data"
	udeli "backend/features/User/delivery"
	uc "backend/features/User/usecase"

	cd "backend/features/Comment/data"
	cdeli "backend/features/Comment/delivery"
	cc "backend/features/Comment/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB, awsConn *session.Session) {
	validator := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, validator)
	userHandler := udeli.New(userCase)
	udeli.RouteUser(e, userHandler)

	postData := pd.New(db)
	postCase := pc.New(postData, validator)
	postHandler := pdeli.New(postCase, awsConn)
	pdeli.RoutePost(e, postHandler)

	commentData := cd.New(db)
	commentCase := cc.New(commentData, validator)
	commentHandler := cdeli.New(commentCase)
	cdeli.RouteComment(e, commentHandler)
}
