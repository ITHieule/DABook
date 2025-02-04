package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

// FILE - ADMIN ROUTER
func RegisterBookRouter(router *gin.RouterGroup) {
	//api - Login
	router.POST("/login", controllers.Book.Login)

	router.POST("/register", controllers.Book.Register)

}
