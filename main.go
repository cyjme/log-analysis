package main

import (
	"github.com/gin-gonic/gin"
	"ideaparLog/database"
	"ideaparLog/controllers"
)

func main() {
	database.Init()
	r := gin.Default()

	goalController := new(controllers.GoalController)
	pvController := new(controllers.PvController)

	r.GET("/goals", goalController.Index)
	r.POST("/goals", goalController.Create)
	r.PUT("/goals/:id", goalController.Update)
	r.DELETE("/goals/:id", goalController.Delete)

	r.GET("/pvs", pvController.List)

	r.GET("/mock/num", func(c *gin.Context) {
		c.JSON(200, gin.H{"num": 100})
	})

	r.Run("0.0.0.0:9999")
}
