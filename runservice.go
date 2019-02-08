package main

import "github.com/gin-gonic/gin"

func responseMsg() string {
	return "pong"
}

func runservice() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": responseMsg(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
