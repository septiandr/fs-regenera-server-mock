package handler

import (
	"net/http"

	"fs-regenera/src/model"
	"fs-regenera/src/services"
	"fs-regenera/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateBookingHandler(c *gin.Context) {
	var req model.BookingRequest

	// Bind & validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"Invalid request data",
			err,
		)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success create data",
		nil, // data
		nil, // meta
	)
}

func GetBookingSummaryHandler(c *gin.Context) {
	var summary model.BookingSummaryResponse

	err := utils.ReadJSONFile("src/data/booking_summarize.json", &summary)
	if err != nil {
		utils.Fail(
			c,
			http.StatusInternalServerError,
			"Failed read booking summary data",
			err,
		)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success get booking summary",
		summary,
		nil,
	)
}

func GetListBookingHandler(c *gin.Context) {
	var query model.BookingListQuery

	// Bind query params
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"Invalid query params",
			err,
		)
		return
	}

	data, total, err := services.GetListBookingService(query)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Failed get booking list", err)
		return
	}

	meta := gin.H{
		"page":  query.Page,
		"limit": query.Limit,
		"total": total,
	}

	utils.Success(c, http.StatusOK, "Success get booking list", data, meta)

}

func GetBookingListLogHandler(c *gin.Context) {
	var query model.BookingDetailQueryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"Invalid query params",
			err,
		)
		return
	}

	// minimal validation
	if query.BookingUUID == "" {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"uuid or code is required",
			nil,
		)
		return
	}

	data, err := services.GetListLogBookingService()
	if err != nil {
		utils.Fail(
			c,
			http.StatusNotFound,
			"Booking not found",
			err,
		)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success validate data",
		data,
		nil,
	)
}
func GetDetailBookingHandler(c *gin.Context) {
	var query model.BookingDetailQueryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"Invalid query params",
			err,
		)
		return
	}

	// minimal validation
	if query.BookingUUID == "" {
		utils.Fail(
			c,
			http.StatusBadRequest,
			"uuid or code is required",
			nil,
		)
		return
	}

	data, err := services.GetDetailBookingService()
	if err != nil {
		utils.Fail(
			c,
			http.StatusNotFound,
			"Booking not found",
			err,
		)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success validate data",
		data,
		nil,
	)
}

func GetBookingByUUIDHandler(c *gin.Context) {
	bookingUUID := c.Param("booking_uuid")
	if bookingUUID == "" {
		utils.Fail(c, http.StatusBadRequest, "booking_uuid is required", nil)
		return
	}

	data, err := services.GetBookingByUUIDService()
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Failed get booking", err)
		return
	}

	utils.Success(c, http.StatusOK, "Success get booking", data, nil)
}
