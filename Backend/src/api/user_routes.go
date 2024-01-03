package api

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	userGroup.GET("/", getUsers)

}

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
