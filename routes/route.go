package routes

import (
	"muslim_pro/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/captions", controller.GetAllCaption)
	r.GET("/captions/:id", controller.GetOneCaption)
	r.GET("captions/random", controller.GetRandomCaption)
	r.POST("/captions", controller.CreateCaption)
	r.PUT("captions/:id", controller.UpdateStatusCaption)
	r.DELETE("/captions/:id", controller.DeleteCaption)

	return r

}
