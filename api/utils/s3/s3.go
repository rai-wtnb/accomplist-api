package s3

import (
	"os"
	"log"
	"fmt"
	"time"
	"errors"
	"bytes"
	"crypto/sha1"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(file []byte, fileName string) (string, error) {
	loadEnv()

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
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		ContentType: aws.String(contentType),
		Key:         aws.String(key),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://s3-%s.amazonaws.com/%s/%s", os.Getenv("AWS_REGION"), os.Getenv("AWS_BUCKET"), key)

	return fileURL, nil
}

// loadEnv prepare to use env.
func loadEnv() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error: failed to load .env file")
    }
}

func uploader() *s3manager.Uploader {
	return s3manager.NewUploader(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     os.Getenv("AWS_ACCESS"),
				SecretAccessKey: os.Getenv("AWS_SECRET"),
			}),
			Region: aws.String(os.Getenv("AWS_REGION")),
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
