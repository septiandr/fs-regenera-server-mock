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

	// TODO: call service layer
	// err := bookingService.Create(req)
	// if err != nil {
	// 	utils.Fail(c, http.StatusInternalServerError, "Failed create booking", err)
	// 	return
	// }

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
