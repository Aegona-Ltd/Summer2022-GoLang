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
		r.GET("form/name/:name", controller.GetBYName)
		r.GET("form/email", controller.GetOrderEmail)
		r.GET("form/date", controller.GetOrderCreateTime)
		r.PUT("form/:id", controller.PutData)
	}
	return r
}
