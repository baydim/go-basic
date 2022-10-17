package controllers

import (
	"go-basic/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFish(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{"message": "pong", "id": id})
}

func GetFishQ(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "pong", "id": id})
}

func PostFish(c *gin.Context) {
	var fish models.Fish
	c.BindJSON(&fish)
	c.JSON(http.StatusOK, gin.H{"message": "pong", "fish": fish})
}
