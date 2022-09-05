package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,

			gin.H{
				"message": "ping pong",
			},
		)
	})
	router.GET("/fish", getFish)
	router.GET("/fish/:id", getFishQ)
	router.POST("/fish", postFish)

	router.Run(":8080")
}

func getFish(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{"message": "pong", "id": id})
}

func getFishQ(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "pong", "id": id})
}

type FishModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func postFish(c *gin.Context) {
	var fish FishModel
	c.BindJSON(&fish)
	c.JSON(http.StatusOK, gin.H{"message": "pong", "fish": fish})

}
