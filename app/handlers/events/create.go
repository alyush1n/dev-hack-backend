package events

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/db/storage"
	"dev-hack-backend/app/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

func Create(c *gin.Context) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Fatal("get file" + err.Error())
	}

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(storage.S3_REGION),
		Credentials: credentials.NewStaticCredentials(config.S3ID, config.S3Secret, ""),
	})
	if err != nil {
		log.Fatal("sessionn: " + err.Error())
	}

	f, err := fileHeader.Open()
	if err != nil {
		log.Fatal("Open header" + err.Error())
	}
	fileName := strings.Replace(fileHeader.Filename, " ", "_", -1)
	url, err := storage.AddFilesToS3(s, fileName, f)

	//url := "https://%s.s3.%s.amazonaws.com/%s"
	//url = fmt.Sprintf(url, storage.S3_BUCKET, storage.S3_REGION, fileName)
	fmt.Println(url)
	jsonInput := struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Date     string `json:"date"`
		SentBy   string `json:"sent_by"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	eventID := primitive.NewObjectID()
	event := model.Event{
		Id:       eventID,
		Type:     jsonInput.Type,
		Name:     jsonInput.Name,
		Location: jsonInput.Location,
		Date:     jsonInput.Date,
		Attachment: model.Attachment{
			Id:     primitive.NewObjectID(),
			URL:    url,
			SentBy: jsonInput.SentBy,
		},
	}

	err = db.InsertEvent(event)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
	})

}
