package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s3Client "presigned.go/services/s3"
)

func GeneratePresignedDownloadUrl(c *gin.Context) {
	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"presignedDownloadUrl": false,
		})
	}

	_, presignClient := s3Client.CreateClient()
	presignedUrl := s3Client.GeneratePresignedDownloadLink(presignClient, reqBody.FileName)
	c.JSON(http.StatusOK, gin.H{
		"presignedUrl": presignedUrl.URL,
	})

}
