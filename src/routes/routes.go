package routes

import (
	"fs-regenera/src/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	// base api
	api := r.Group("/api/v1/cms")

	// ================= AUTH =================
	api.POST("/login", handler.LoginHandler)
	api.GET("/profiles", handler.ProfileHandler)

	// ================= MERCHANT =================
	api.GET("/merchants", handler.GetMerchantsListHanlder)
	api.GET("/merchants/:merchant_uuid/outlets", handler.GetOutletsListHandler)

	// ================= DOCTOR =================
	api.GET("/doctors", handler.GetDoctorsListHandler)
	api.GET("/doctors/:doctor_uuid/sessions", handler.GetDoctorSessions)
	api.GET("/doctors/:doctor_uuid/booked/:date/count", handler.GetListDoctorBookedHandler)

	// ================= CUSTOMER =================
	api.GET("/customers", handler.GetListCustomersHandler)

	// ================= OUTLETS =================
	api.GET("/outlets/:outlet_uuid", handler.GetOutletDetailHandler)

	// ================= BOOKING =================
	api.POST("/bookings", handler.CreateBookingHandler)
	api.GET("/bookings/summaries", handler.CreateBookingHandler)
	api.GET("/bookings", handler.GetListBookingHandler)
	api.GET("/bookings/:booking_uuid/logs", handler.GetBookingListLogHandler)
	api.GET("/bookings/:booking_uuid", handler.GetBookingByUUIDHandler)
}
