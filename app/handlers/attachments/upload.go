package attachments

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Upload(c *gin.Context) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Fatal("get file" + err.Error())
	}

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(storage.S3_REGION),
		Credentials: credentials.NewStaticCredentials(config.S3ID, config.S3Secret, ""),
	})
	if err != nil {
		log.Fatal("session	: " + err.Error())
	}

	f, err := fileHeader.Open()
	if err != nil {
		log.Fatal("Open header" + err.Error())
	}
	fileName := strings.ReplaceAll(fileHeader.Filename, " ", "_")
	url, err := storage.AddFilesToS3(s, fileName, f)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"url":     url,
	})
}
