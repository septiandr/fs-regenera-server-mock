package main

import (
	"fs-regenera/src/middleware"
	"fs-regenera/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.ApiIDMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	routes.Routes(r)

	r.Run(":9070")
}
