package main

import (
	"fs-regenera/src/handler"
	"fs-regenera/src/middleware"

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

	behaveGroup := r.Group("/v1/cms")
	{
		behaveGroup.GET("/merchants/:merchant_uuid/outlets", handler.GetOutletsListHandler)
	}

	r.Run(":9070")
}
