package s3Client

import (
	"context"
	"log"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"presigned.go/constants"
)

func GeneratePresignedUrl(presignClient *s3.PresignClient, fileName string, fileSize int64) *v4.PresignedHTTPRequest {
	bucketName := constants.BUCKET_NAME
	objectKey := fileName
	lifetimeSecs := int64(600)

	request, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(objectKey),
		ContentLength: aws.Int64(fileSize),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}

	return request

}
