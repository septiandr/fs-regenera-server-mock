package routes

import (
	"fs-regenera/src/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	behaveGroup := r.Group("/api/v1/cms")
	{
		behaveGroup.GET("/merchants/:merchant_uuid/outlets", handler.GetOutletsListHandler)
		behaveGroup.GET("/doctors", handler.GetDoctorsListHandler)
		behaveGroup.GET("/doctors/:doctor_uuid/sessions", handler.GetDoctorSessions)
	}
}
