package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {

	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
