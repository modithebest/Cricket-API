package routes

import (
	"cricket_stats_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/hello", controllers.HealthCheck)
	router.POST("/stats", controllers.Stats)
	router.GET("/Stats", controllers.GetStats)
	router.GET("/statistics", controllers.Statisitic)
	router.POST("/statisticsPost", controllers.StatPost)
	router.POST("/db", controllers.Post)

}
