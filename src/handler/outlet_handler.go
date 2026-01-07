package handler

import (
	"fs-regenera/src/middleware"
	"fs-regenera/src/model"
	"fs-regenera/src/services"
	"fs-regenera/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOutletsListHandler(c *gin.Context) {
	merchantUUID := c.Param("merchant_uuid")
	_ = merchantUUID // currently not used
	var params model.OutletListParams

	//validasi query params
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.Fail(c, http.StatusBadRequest, "Invalid request params", err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	data, total, error := services.GetOutletListServices(c.Request.Context(), params)
	if error != nil {
		utils.Fail(c, http.StatusInternalServerError, "Failed to get outlet list", error)
		return
	}

	utils.Success(c, http.StatusOK, "Outlet list retrieved successfully", data, middleware.MetaPagination{
		Page:       params.Page,
		Limit:      params.Limit,
		Total:      total,
		TotalPages: (total + params.Limit - 1) / params.Limit,
	})
}
