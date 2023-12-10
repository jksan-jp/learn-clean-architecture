package infrastructure

import (
	"learn-clean-architecture/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userContoroller := controllers.NewUserController(NewSqlHander())

	router.POST("/users", func(c *gin.Context) { userContoroller.Create(c) })

	Router = router
}
