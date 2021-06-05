package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	S3_REGION = "us-east-2"
	S3_BUCKET = "dev-hack-backend"
)

func AddFilesToS3(s *session.Session, fileName string, r io.Reader) (string, error) {

	body, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, len(body))
	b := bytes.NewReader(body)
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 b,
		ContentLength:        aws.Int64(int64(len(body))),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		fmt.Println(err)
	}

	req, _ := s3.New(s).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fileName),
	})

	url, err := req.Presign(time.Hour * 24)
	if err != nil {
		return "", fmt.Errorf("failed to presign GetObjectRequest for key %q: %v", fileName, err)
	}
	return url, nil
}
