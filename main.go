package main

import (
	ctl "github.com/donghquinn/go_web/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/", ctl.SayHelloController)
	router.POST("/test", ctl.PostTestController)
	router.Run()
}