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

	router.GET("/codebreacker/setup/:number", func(c *gin.Context){
		number := c.Param("number")
		c.String(http.StatusOK, "secret number configured: " + number)
	})

	router.GET("/codebreacker/guess/:number", func(c *gin.Context) {
		name := c.Param("number")
		result := validate(name)
		c.String(http.StatusOK, "Answer: ", result)
	})

	router.Run(":" + port)
}