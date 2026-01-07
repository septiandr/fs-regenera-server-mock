package main

import (
	"fs-regenera/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Use(middleware.ApiIDMiddleware())

	behaveGroup := r.Group("/v1/cms")
	{
		behaveGroup.GET("/merchants/{merchant_uuid}/outlets", getOutletsListHandler)
	}

	r.Run()
}
