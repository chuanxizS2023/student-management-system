package main

import (
	"SMS/src/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "from main page",
		})
	})
	api.UserRoutes(r)

	r.Run(":9000")
}
