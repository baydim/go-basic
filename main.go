package main

import (
	"go-basic/controllers"

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
	router.GET("/fish", controllers.GetFish)
	router.GET("/fish/:id", controllers.GetFishQ)
	router.POST("/fish", controllers.PostFish)
	router.POST("/upload", controllers.UploadSingleFile)
	router.GET("/delete/:filename", controllers.DeleteSingleFile)
	router.GET("/fetch/:id", controllers.FetchApi)
	router.GET("/fetchModel/:id", controllers.FetchApiUsingModel)
	router.GET("/scrape", controllers.ScrapingWeb)
	router.Run(":1234")
}
