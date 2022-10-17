package controllers

import (
	"encoding/json"
	"go-basic/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnmarshalModelTodo(data []byte) (models.ModelTodo, error) {
	var r models.ModelTodo
	err := json.Unmarshal(data, &r)
	return r, err
}

func FetchApiUsingModel(c *gin.Context) {
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
	var model models.ModelTodo
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
