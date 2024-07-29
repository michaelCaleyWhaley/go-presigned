package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s3Client "presigned.go/services/s3"
)

func Home(c *gin.Context) {

	client, _ := s3Client.CreateClient()
	objectList := s3Client.ListObjects(client)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":      "Presigned",
		"objectList": objectList,
	})
}
