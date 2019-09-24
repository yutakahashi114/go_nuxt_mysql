package infrastructure

import (
	"app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	api := router.Group("/api")
	{
		api.POST("/users", func(c *gin.Context) { userController.Create(c) })
		api.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
		api.GET("/all-users", func(c *gin.Context) { userController.All(c) })
	}
	Router = router
}
