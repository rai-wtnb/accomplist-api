package s3

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(file []byte, fileName string) (string, error) {
	fileUrl, err := sendToS3(file, fileName)
	if err != nil {
		return fileUrl, err
	}
	return fileUrl, nil
}

func sendToS3(file []byte, fileName string) (string, error) {
	extension := filepath.Ext(fileName)
	contentType := getContentType(extension)
	if contentType == "" {
		log.Println("contentType: ", contentType)
		return "", errors.New("Unknown content type")
	}

	uploader := uploader()

	h := sha1.Sum(file)
	key := fmt.Sprintf("uploads/%s_%x", time.Now().Format("20060102150405"), h[:4])

	_, err := uploader.Upload(&s3manager.UploadInput{
		Body:        bytes.NewReader(file),
		Bucket:      aws.String("accomplist-bucket"),
		ContentType: aws.String(contentType),
		Key:         aws.String(key),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://s3-ap-northeast-1.amazonaws.com/accomplist-bucket/%s", key)

	return fileURL, nil
}

func uploader() *s3manager.Uploader {
	return s3manager.NewUploader(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     os.Getenv("S3_ACCESS"),
				SecretAccessKey: os.Getenv("S3_SECRET"),
			}),
			Region: aws.String("ap-northeast-1"),
		},
	})))
}

// getContentType judges file file type.
func getContentType(extension string) string {
	switch extension {
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".png":
		return "image/png"
	default:
		return ""
	}
}
