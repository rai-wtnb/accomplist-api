package s3

import (
	"os"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(file string) {
	loadEnv()
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS"), os.GetEnv("AWS_SECRET"), ""),
	  Region: aws.String("ap-northeast-1"),
	 }))
}

func Get() {
	log.Println("get from s3")
}

// loadEnv prepare to use env.
func loadEnv() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error: failed to load .env file")
    }
}
