package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, " Integración con CodeShip, Por: Vanessa Pérez")
	})

	router.GET("/codebreaker/setup/:number", func(c *gin.Context) {
		number := c.Param("number")
		setSecret(number)
		c.String(http.StatusOK, "secret number configured: "+number)
	})

	router.GET("/codebreaker/guess/:number", func(c *gin.Context) {
		name := c.Param("number")
		result := guess(name)
		c.String(http.StatusOK, "Answer: %s", result)
	})

	router.Run(":" + port)
}
