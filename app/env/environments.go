package env

import (
	"os"
)

var (
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	S3AKID           string
	S3SecretKey      string
	AwsS3Bucket      string
)

func Init() {
	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresDB = os.Getenv("POSTGRES_DB")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	S3AKID = os.Getenv("S3_AKID")
	S3SecretKey = os.Getenv("S3_SECRET_KEY")
	AwsS3Bucket = os.Getenv("Aws_S3_Bucket")
}
