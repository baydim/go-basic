package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadSingleFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, file.Filename)
	filename := filepath.Base(file.Filename)
	response := c.SaveUploadedFile(file, filename)
	if response == nil {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "filename": filename})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "filename": filename})
	}
}

func DeleteSingleFile(c *gin.Context) {
	response := os.Remove(c.Param("filename"))
	if response == nil {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "filename": c.Param("filename")})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "filename": c.Param("filename")})
	}
}
