package module

import (
	ctl "github.com/donghquinn/go_web/controllers"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	router.GET("/", ctl.SayHelloController)
	router.POST("/test", ctl.PostTestController)
	router.POST("/sum", ctl.CalculateSumController)
	router.POST("/add", ctl.CalculateAddController)
}