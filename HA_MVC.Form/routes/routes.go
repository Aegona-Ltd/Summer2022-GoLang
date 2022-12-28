package routes

import (
	"HA_MVC.Form/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	{
		r.POST("form", controller.PostForm)
		r.GET("form", controller.GetALLForm)
		r.GET("form/:email", controller.GetBYEmail)
		r.GET("form/:email", controller.GetBYName)
	}
	return r
}
