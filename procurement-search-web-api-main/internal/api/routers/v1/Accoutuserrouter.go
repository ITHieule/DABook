package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccoutRouter(router *gin.RouterGroup) {
	router.POST("/add", controllers.Procurement.AccoutUser)
	router.POST("/Search", controllers.Procurement.SearchUser)
	router.POST("/Delete", controllers.Procurement.DeleteUser)
	router.POST("/Update", controllers.Procurement.UpdateUser)

}
