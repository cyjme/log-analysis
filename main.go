package main

import (
	"github.com/gin-gonic/gin"
	"ideaparLog/database"
	"ideaparLog/controllers"
	"fmt"
	"time"
	"ideaparLog/cmd/readPvLog"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(CORSMiddleware())

	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			fmt.Println("start read log")
			readPvLog.Run()
		}
	}()

	goalController := new(controllers.GoalController)
	pvController := new(controllers.PvController)

	r.GET("/goals", goalController.Index)
	r.POST("/goals", goalController.Create)
	r.PUT("/goals/:id", goalController.Update)
	r.DELETE("/goals/:id", goalController.Delete)

	r.GET("/pvs", pvController.List)
	r.GET("/countQuery", pvController.CountQuery)

	r.GET("/mock/num", func(c *gin.Context) {
		c.JSON(200, "20")
	})

	r.Run("0.0.0.0:9999")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
