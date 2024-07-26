package main

import (
	"github.com/gin-gonic/gin"
	controllerPresigned "presigned.go/controllers/generate-upload-url"
	controllerHome "presigned.go/controllers/home"
	"presigned.go/initialisers"
)

func main() {
	initialisers.LoadEnvVars()

	// https: //docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllerHome.Home)
	r.POST("/presigned", controllerPresigned.GeneratePresignedUrl)
	r.Run("localhost:3000")
}
