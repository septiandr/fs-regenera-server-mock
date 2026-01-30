package handler

import (
	"fs-regenera/src/model"
	"fs-regenera/src/services"
	"fs-regenera/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListCustomersHandler(c *gin.Context) {
	var query model.CustomerCheckQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"Invalid query params",
			err,
		)
		return
	}

	// ✅ DEFAULT VALUE (INI KUNCI)
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	data, total, err := services.GetListCustomersService(query)
	if err != nil {
		utils.Fail(
			c,
			http.StatusInternalServerError,
			"Failed get customers",
			err,
		)
		return
	}

	meta := gin.H{
		"page":  query.Page,
		"limit": query.Limit,
		"total": total,
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success get customer list",
		data,
		meta,
	)
}
