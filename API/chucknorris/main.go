package main

import (
	"chucknorris/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/api/random-fact", func(c *gin.Context) {

		fact := Controller.RandomFact()
		c.JSON(http.StatusOK, gin.H{
			"fato": fact,
		})
		//fmt.Printf("The random fact is %s", fact)
	})

	router.Run()
}
