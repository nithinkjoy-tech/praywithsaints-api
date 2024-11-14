package routes

import (
	"prayer-book/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/prayer", controllers.GetPrayers)
	// router.GET("/prayer/:id", controllers.GetPrayerById)
	router.POST("/prayer", controllers.InsertPrayer)

	return router
}
