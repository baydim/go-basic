package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	router.POST("/upload", uploadSingleFile)
	router.GET("/delete/:filename", deleteSingleFile)
	router.GET("/fetch/:id", fetchApi)
	router.Run(":1234")
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

func uploadSingleFile(c *gin.Context) {
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

func deleteSingleFile(c *gin.Context) {
	response := os.Remove(c.Param("filename"))
	if response == nil {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "filename": c.Param("filename")})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "filename": c.Param("filename")})
	}
}

func fetchApi(c *gin.Context) {
	id := c.Params.ByName("id")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "error": err})
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "error": err})
	}
	//Convert the body to type string
	sb := json.RawMessage(body)

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "pong", "body": sb})
}
