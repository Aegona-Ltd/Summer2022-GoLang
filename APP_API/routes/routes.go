package routes

import (
	"APP_API/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())
	{
		// User
		r.POST("user", controller.Create_User)
		r.PUT("user/:id", controller.Update_User)

		//Admin
		r.GET("user", controller.GetALL_User)
		r.GET("user/:id", controller.GetById_User)
		r.DELETE("user/:id", controller.Delete_User)
		r.PUT("role_user/:id", controller.Update_Role_User)

		// r.GET("role_user", controller.GetALL_Role_User)
	}
	return r
}
