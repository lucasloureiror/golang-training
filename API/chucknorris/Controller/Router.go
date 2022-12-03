package Controller

import (
	"chucknorris/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get() {

	router := gin.Default()

	router.GET("/api/random-fact", func(c *gin.Context) {
		var fact model.Fact
		fact = RandomFact()
		c.JSON(http.StatusOK, gin.H{
			"id":    fact.Id,
			"value": fact.Value,
		})
	})

	router.GET("/api/categories", func(c *gin.Context) {
		categories := Categorias()
		c.String(http.StatusOK, categories)
	})

	router.GET("/api/categories/search/:category", func(c *gin.Context) {
		search := c.Param("category")
		fact := RandomFactWithCategory(search)
		c.JSON(http.StatusOK, gin.H{
			"id":    fact.Id,
			"value": fact.Value,
		})
	})

	router.Run()
}

func post() {

}
