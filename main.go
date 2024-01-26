package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/donghquinn/go_web/module"
	"github.com/donghquinn/go_web/utilities"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	// Load env
	godotenv.Load(".env")

	// Set variable
	port := os.Getenv("APP_PORT")
	mode := os.Getenv("ENVIRONMENT")

	// Server Mode
	gin.SetMode(mode)

	router := gin.Default()

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Printf("Server Listening On %s\n", port)
	fmt.Printf("Server Mode: %s\n", mode)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")

	module.Handler(router)

	// Graceful ShutDown
	utilities.GracefulShutDown(server)
}