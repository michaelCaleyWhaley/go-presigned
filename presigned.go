package main

import (
	"net/http"

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
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "Login",
		})
	})

	r.POST("/presigned", controllers.GeneratePresignedUrl)
	r.POST("/download", controllers.GeneratePresignedDownloadUrl)
	r.Run("localhost:3000")
}
