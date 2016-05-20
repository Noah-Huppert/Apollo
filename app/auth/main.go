package main

import "github.com/gin-gonic/gin"
import "Apollo/app/auth/http/resp"

func root(c *gin.Context) {
	resp.Resp(c)
}

func main() {
	router := gin.Default()

	router.GET("/", root) 

	router.Run(":3000")
}
