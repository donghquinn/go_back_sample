package main

import (
	"net/http"

	"github.com/donghquinn/go_web/module"
	"github.com/donghquinn/go_web/utilities"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	module.Handler(router)
	
	server := &http.Server{
		Addr:    ":8594",
		Handler: router,
	}

	// Graceful ShutDown
	utilities.GracefulShutDown(server)
}