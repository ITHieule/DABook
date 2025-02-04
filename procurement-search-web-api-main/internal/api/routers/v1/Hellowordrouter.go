package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHelloWordRouter(router *gin.RouterGroup) {
	router.GET("/get", controllers.Procurement.GetHelloWorld)

}
