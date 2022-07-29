package aws

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
)

func InitS3(key, secret, region string) *session.Session {
	conn, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				key, secret, "",
			),
		},
	)
	if err != nil {
		log.Error("S3 Config error:", err)
	}
	return conn
}

func DoUpload(sess *session.Session, file multipart.FileHeader, filename string) string {
	manager := s3manager.NewUploader(sess)
	src, err := file.Open()
	if err != nil {
		log.Info(err)
	}
	defer src.Close()
	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()

	res, err := manager.Upload(
		&s3manager.UploadInput{
			Bucket:      aws.String("be10-file"),
			ACL:         aws.String("public-read"),
			ContentType: aws.String(http.DetectContentType(buffer)),
			Key:         aws.String(filename),
			Body:        body,
		},
	)
	if err != nil {
		log.Info(res)
		log.Error("Upload error : ", err)
	}
	link := fmt.Sprint("https://be10-file.s3.ap-southeast-1.amazonaws.com/", filename)
	return link
}
