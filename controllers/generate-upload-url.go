package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"presigned.go/constants"
	s3Client "presigned.go/services/s3"
)

type ReqBody struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

func GeneratePresignedUrl(c *gin.Context) {
	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"presignedUrl": false,
		})
		return
	}

	if reqBody.FileSize > constants.BYTES_GIGABYTE {
		c.JSON(http.StatusBadRequest, gin.H{
			"presignedUrl": false,
			"message":      "Exceeded file size limit.",
		})
		return
	}

	_, presignClient := s3Client.CreateClient()
	presignedUrl := s3Client.GeneratePresignedUrl(presignClient, reqBody.FileName, reqBody.FileSize)
	c.JSON(http.StatusOK, gin.H{
		"presignedUrl": presignedUrl.URL,
	})
	return

}
