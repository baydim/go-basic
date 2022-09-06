package main

import (
	"encoding/json"
	"io/ioutil"

	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
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
	router.GET("/fetchModel/:id", fetchApiUsingModel)
	router.GET("/scrape", scrapingWeb)
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

func UnmarshalModelTodo(data []byte) (ModelTodo, error) {
	var r ModelTodo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ModelTodo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ModelTodo struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchApiUsingModel(c *gin.Context) {
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
	var model ModelTodo
	err = json.Unmarshal(sb, &model)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "pong", "error": err})
	}
	if model.Completed {
		c.JSON(http.StatusOK, gin.H{
			"complete": true,
			"id":       id,
			"message":  "pong", "body": model})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"complete": false,
			"id":       id,
			"message":  "pong", "body": model})
	}

}

func scrapingWeb(c *gin.Context) {
	doc, err := goquery.NewDocument("https://baydim.github.io/yournose/")
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong", "title": title})

}
