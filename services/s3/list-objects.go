package s3Client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"presigned.go/constants"
)

type ContentObj struct {
	Key  string
	Size int64
}

func ListObjects(client *s3.Client) []ContentObj {

	maxKeys := int32(20)
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: &constants.BUCKET_NAME, MaxKeys: &maxKeys})

	if err != nil {
		log.Printf("Couldn't get a list of objects from %v. Here's why: %v\n",
			&constants.BUCKET_NAME, err)
	}

	var contentList []ContentObj

	for _, content := range output.Contents {
		contentList = append(contentList, ContentObj{Key: *content.Key, Size: *content.Size})
	}

	return contentList

}
