package s3Client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateClient() (*s3.Client, *s3.PresignClient) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	return client, presignClient
}
