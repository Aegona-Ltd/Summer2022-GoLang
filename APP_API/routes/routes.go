package routes

import (
	"APP_API/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	{
		r.POST("user", controller.Create_User)
		r.GET("user", controller.GetALL_User)
		r.GET("user/:id", controller.GetById_User)
		r.PUT("user/:id", controller.Update_User)
		r.DELETE("user/:id", controller.Delete_User)

		r.POST("role", controller.Create_Role)
		r.GET("role", controller.GetALL_Role)
		r.GET("role/:id", controller.GetById_Role)
		r.PUT("role/:id", controller.Update_Role)
		r.DELETE("role/:id", controller.Delete_Role)
	}
	return r
}
