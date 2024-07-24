package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"presigned.go/initialisers"
	s3Client "presigned.go/services/s3"
)

func main() {
	initialisers.LoadEnvVars()

	client, presignClient := s3Client.CreateClient()
	s3Client.LogS3Buckets(client)

	bucketName := "caleycodelab.com"
	objectKey := "test_pic.png"
	lifetimeSecs := int64(600)

	// https://medium.com/@aidan.hallett/securing-aws-s3-uploads-using-presigned-urls-aa821c13ae8d
	// create bucket that can allow puts for testing

	request, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}

	log.Println("request", request)

	// https: //docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html

	// r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/", controllers.Home)
	// r.Run(":3000")
}
