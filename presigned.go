package main

import (
	"github.com/gin-gonic/gin"
	controllers "presigned.go/controllers/home"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controllers.Home)

	r.Run(":3000")
}
