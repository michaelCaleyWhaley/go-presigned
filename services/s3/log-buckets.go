package s3Client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LogS3Buckets(client *s3.Client) {
	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, bucket := range output.Buckets {
		log.Println(*bucket.Name)
	}
}
