package handler

import (
	"fs-regenera/src/model"
	"fs-regenera/src/services"
	"fs-regenera/src/utils"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, 400, "Invalid request body", err)
		return
	}

	token, err := services.LoginService(c.Request.Context(), req)
	if err != nil {
		utils.Fail(c, 401, "Email or password incorrect", err)
		return
	}

	utils.Success(c, 200, "Login successful", map[string]string{
		"access_token": token,
	}, nil)
}
