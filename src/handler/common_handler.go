package handler

import (
	"fs-regenera/src/utils"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {

	res := map[string]interface{}{
		"uuid":       "c28f392b-337c-41eb-8277-03bc0246d89b",
		"name":       "Super Admin",
		"email":      "superadmin@mailinator.com",
		"image":      "",
		"is_pin_set": false,
		"created_at": "2025-12-18T02:02:51.367231Z",
		"updated_at": "2025-12-18T02:02:51.367231Z",
	}

	utils.Success(c, 200, "Profile fetched successfully", res, nil)

}
