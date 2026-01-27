package handler

import (
	"net/http"

	"fs-regenera/src/model"
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
