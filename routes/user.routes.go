package routes

import (
	"golang-gin-gonic-studing/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    
    r.GET("/users", controller.GetUsers)
    r.POST("/users", controller.CreateUser)

    return r
}
