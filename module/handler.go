package module

import (
	ctl "github.com/donghquinn/go_web/controllers"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	router.GET("/", ctl.SayHelloController)
	router.POST("/test", ctl.PostTestController)

	// Grouping
	calc := router.Group("/calculate")

	{
		calc.POST("/sum", ctl.CalculateSumController)
		calc.POST("/add", ctl.CalculateAddController)
	}

	user := router.Group("/client")
	{
		user.POST("/total", ctl.UserCountControllers)
	}
}
