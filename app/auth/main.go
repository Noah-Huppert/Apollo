package main

import "github.com/gin-gonic/gin"

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"key": "value",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", root) 

	router.Run(":3000")
}
