package main

import (
	"SMS/src/api"

	"github.com/gin-gonic/gin"

	"SMS/src/internal/database"

	"SMS/src/configs"
)

func main() {
	r := gin.Default()

	//config project
	configs.LoadEnv()

	database.InitDB()

	defer database.GetDB().Close()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "from main page",
		})
	})
	api.UserRoutes(r)

	r.Run(":9090")
}
