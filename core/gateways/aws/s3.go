package aws

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

type S3 struct {
	Bucket  string
	session *session.Session
}

func NewS3(bucket string) *S3 {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Println("AWS S3 error", err)
	}
	return &S3{
		session: s,
		Bucket:  bucket,
	}
}

func (h *S3) Upload(name string, body []byte, contentType string) error {
	_, err := s3.New(h.session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(h.Bucket),
		Key:                  aws.String(name),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(body),
		ContentLength:        aws.Int64(int64(len(body))),
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		log.WithField("error message", err.Error()).Error("Failed to upload file to bucket")
	}
	return err
}

func (h *S3) Delete(name string) error {
	name, err := url.QueryUnescape(name)
	if err != nil {
		log.Println("[AWS GET LINK]:", err)
	}

	_, err = s3.New(h.session).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(name),
	})
	return err
}

func (h *S3) GetPublicSignedPathForObject(name string) string {
	name, err := url.QueryUnescape(name)
	if err != nil {
		log.Println("[AWS GET LINK]:", err)
	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(name),
	}

	req, _ := s3.New(h.session).GetObjectRequest(params)

	url, err := req.Presign(1 * time.Minute) // Set link expiration time
	if err != nil {
		log.Println("[AWS GET LINK]:", params, err)
	}

	return url
}

func (h *S3) GetPublicPathForObject(name string) string {
	url := "https://%s.s3.%s.amazonaws.com/%s"
	return fmt.Sprintf(url, os.Getenv("AWS_BUCKET"), os.Getenv("AWS_REGION"), name)
}

func (h *S3) GetPutUrl(name, mimeType string) (string, error) {

	params := &s3.PutObjectInput{
		Bucket:      aws.String(h.Bucket),
		ACL:         aws.String("public-read"),
		Key:         aws.String(name),
		ContentType: aws.String(mimeType),
	}

	req, _ := s3.New(h.session).PutObjectRequest(params)

	url, err := req.Presign(30 * time.Minute) // Set link expiration time
	if err != nil {
		log.Println("[AWS GET LINK]:", params, err)
	}

	return url, err
}
