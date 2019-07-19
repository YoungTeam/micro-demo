package router

import (
	"mgo/api-srv/controller"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	//router := gin.New()
	// router.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	AllowCredentials: true,
	// }))
	// api := router.Group("/api/v1")
	// {
	// 	api.POST("/user/registry", v1.RegistryUser)
	// 	api.POST("/user/login", v1.UserLogin)
	// }

	router.GET("/user/", controller.GetUser)

}
