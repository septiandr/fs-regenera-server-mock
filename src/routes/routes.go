package routes

import (
	"fs-regenera/src/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	behaveGroup := r.Group("/api/v1/cms")
	{
		behaveGroup.POST("/login", handler.LoginHandler)
		behaveGroup.GET("/profile", handler.ProfileHandler)
		behaveGroup.GET("/merchants/:merchant_uuid/outlets", handler.GetOutletsListHandler)
		behaveGroup.GET("/merchants", handler.GetMerchantsListHanlder)
		behaveGroup.GET("/doctors", handler.GetDoctorsListHandler)
		behaveGroup.GET("/doctors/:doctor_uuid/sessions", handler.GetDoctorSessions)
		behaveGroup.GET("/bookings/:doctor_uuid/booked/:date/count", handler.GetListDoctorBookedHandler)
	}
}
