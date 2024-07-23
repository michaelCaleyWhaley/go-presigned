package main

import (
	"presigned.go/initialisers"
	s3Client "presigned.go/services/s3"
)

func main() {
	initialisers.LoadEnvVars()

	client := s3Client.CreateClient()
	s3Client.LogS3Buckets(client)

	// https: //docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html

	// r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/", controllers.Home)
	// r.Run(":3000")
}
