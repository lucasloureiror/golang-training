package main

import (
	"chucknorris/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/api/random-fact", func(c *gin.Context) {
		var fact Controller.Fact

		fact = Controller.RandomFact()
		c.JSON(http.StatusOK, gin.H{
			"id":    fact.Id,
			"value": fact.Value,
		})
		//fmt.Printf("The random fact is %s", fact)
	})

	router.Run()
}
