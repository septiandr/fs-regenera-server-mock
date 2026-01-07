package middleware

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	ApiID   string      `json:"api_id"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	ApiID   string `json:"api_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []any  `json:"data"`
	Error   string `json:"error,omitempty"`
}

type MetaPagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func ApiIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID := fmt.Sprintf(
			"API_CALL_%d_%d",
			time.Now().UnixMilli(),
			rand.Intn(1_000_000),
		)

		c.Set("api_id", apiID)
		c.Next()
	}
}
