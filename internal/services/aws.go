package services

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(bucket, key, body string) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   aws.ReadSeekCloser(strings.NewReader(body)),
	})
	if err != nil {
		log.Fatal("Failed to upload file to S3:", err)
	}
}
