package main

import (
	"net/http"

	ctl "github.com/donghquinn/go_web/controllers"
	"github.com/donghquinn/go_web/utils"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/", ctl.SayHelloController)
	router.POST("/test", ctl.PostTestController)

	server := &http.Server{
		Addr:    ":8594",
		Handler: router,
	}

	// Graceful ShutDown
	utils.GracefulShutDown(server)
}