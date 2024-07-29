package main

import (
	"github.com/gin-gonic/gin"
	"presigned.go/controllers"
	"presigned.go/initialisers"
)

func main() {
	initialisers.LoadEnvVars()

	// https: //docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Home)
	r.POST("/presigned", controllers.GeneratePresignedUrl)
	r.POST("/download", controllers.GeneratePresignedDownloadUrl)
	r.Run("localhost:3000")
}
