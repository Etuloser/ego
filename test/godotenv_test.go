package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGoDotEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println(s3Bucket, secretKey)
}
