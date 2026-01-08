package main

import (
	"fs-regenera/src/middleware"
	"fs-regenera/src/routes"
	"fs-regenera/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", utils.PasswordValidator)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.ApiIDMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	routes.Routes(r)

	r.Run(":9070")
}
