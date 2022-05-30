package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

func main() {
	// comment this out if you want to work on a real AWS account
	err := os.Setenv("LOCALSTACK_ENDPOINT", "http://localhost:4566")
	if err != nil {
		panic(err)
	}

	// when running locally, we're going to have the LOCALSTACK_ENDPOINT environment variable set
	sess, err := CreateSession("eu-central-1")
	if err != nil {
		panic(err)
	}

	sqsSvc := sqs.New(sess)
	s3Svc := s3manager.NewUploader(sess)

	for {
		result, err := sqsSvc.ReceiveMessage(
			&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String("http://localhost:4566/000000000000/queue1"),
				MaxNumberOfMessages: aws.Int64(1),
				WaitTimeSeconds:     aws.Int64(3),
			},
		)

		if err != nil {
			fmt.Printf("failed to receive message with error %v", err)
			continue
		}

		if len(result.Messages) == 0 {
			continue
		}

		_, err = s3Svc.Upload(&s3manager.UploadInput{Bucket: aws.String("bucket1"),
			Key:  aws.String("test1"),
			Body: bytes.NewReader([]byte(*result.Messages[0].Body))},
		)
		if err != nil {
			fmt.Printf("failed to upload key with error %v", err)
		}
	}
}

func CreateSession(region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	if localStackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT"); localStackEndpoint != "" {
		awsConfig.S3ForcePathStyle = aws.Bool(true)
		awsConfig.Endpoint = aws.String(localStackEndpoint)
	}

	return session.NewSession(awsConfig)
}
