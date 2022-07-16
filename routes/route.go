package routes

import (
	"fmt"
	"github.com/MeGaNeKoS/TF-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
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

func Setup(router *gin.Engine) {
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/get", controllers.GetAnimal)
			v1.GET("/get/:id", controllers.GetAnimalById)
			v1.DELETE("/delete", controllers.DeleteAnimal)
			v1.POST("/add", controllers.InputAnimal)
			v1.PUT("/update", controllers.UpdateAnimal)
		}
	}
}
