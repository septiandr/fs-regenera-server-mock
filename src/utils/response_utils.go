package utils

import (
	"fs-regenera/src/middleware"

	"github.com/gin-gonic/gin"
)

func Fail(
	c *gin.Context,
	httpStatus int,
	message string,
	err error,
) {
	apiID, _ := c.Get("api_id")

	resp := middleware.ErrorResponse{
		ApiID:   apiID.(string),
		Status:  "FAILED",
		Message: message,
		Data:    []any{}, // <-- SELALU []
	}

	if err != nil {
		resp.Error = err.Error()
	}

	c.JSON(httpStatus, resp)
}

func Success(
	c *gin.Context,
	httpStatus int,
	message string,
	data interface{},
	meta interface{},
) {
	apiID, _ := c.Get("api_id")

	c.JSON(httpStatus, middleware.SuccessResponse{
		ApiID:   apiID.(string),
		Status:  "SUCCESS",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
