package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s3Client "presigned.go/services/s3"
)

type ReqBody struct {
	FileName string `json:"fileName"`
}

func GeneratePresignedUrl(c *gin.Context) {
	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"presignedUrl": false,
		})
	}

	_, presignClient := s3Client.CreateClient()
	presignedUrl := s3Client.GeneratePresignedUrl(presignClient, reqBody.FileName)
	c.JSON(http.StatusOK, gin.H{
		"presignedUrl": presignedUrl.URL,
	})

}
