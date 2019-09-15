package infrastructure

import (
	"app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.GET("/users", func(c *gin.Context) { userController.All(c) })
	Router = router
}
