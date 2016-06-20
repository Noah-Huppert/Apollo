package main

import "github.com/gin-gonic/gin"
import "Apollo/app/auth/http/resp"

func routePublicConfig(c *gin.Context) {
	resp.Resp(gin.H { 
		"client_id": "58271217522-1qu3oa4aqh7kcvb2gfs9ouhe6jp1tthi.apps.googleusercontent.com",
		"oauth_callback": "http://127.0.0.1:8080/v1/auth/google/oauth2callback",
	}, 200, c)
}

func main() {
	router := gin.Default()

	router.GET("/v1/auth/google/public_config", routePublicConfig)

	router.Run(":3001")
}
