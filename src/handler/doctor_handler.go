package handler

import (
	"fs-regenera/src/middleware"
	"fs-regenera/src/model"
	"fs-regenera/src/services"
	"fs-regenera/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDoctorsListHandler(c *gin.Context) {
	var params model.DoctorListParams

	if err := c.ShouldBindQuery(&params); err != nil {
		utils.Fail(c, http.StatusBadRequest, "Invalid request params", err)
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	data, total, error := services.GetDoctorListServices(c, params)

	if error != nil {
		utils.Fail(c, http.StatusInternalServerError, "Failed to get doctor list", error)
		return
	}

	utils.Success(c, http.StatusOK, "Doctor list retrieved successfully", data, middleware.MetaPagination{
		Page:       params.Page,
		Limit:      params.Limit,
		Total:      total,
		TotalPages: (total + params.Limit - 1) / params.Limit,
	})

}
