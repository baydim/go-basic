package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func FetchApi(c *gin.Context) {
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

func ScrapingWeb(c *gin.Context) {
	doc, err := goquery.NewDocument("https://baydim.github.io/yournose/")
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong", "title": title})

}
