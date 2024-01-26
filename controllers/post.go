package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostBodyTestForm struct {
	Name string `json:"name"`
	Message string `json:"message"`
}

func PostTestController(ctx *gin.Context) {
	request := PostBodyTestForm{}

	fmt.Println(request)

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"errorMessage": "Required Value is not Included"})
	} else {
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		
		fmt.Printf("Name: %s, Message: %s", name, message)
		ctx.JSON(http.StatusOK, gin.H{"name": name, "message": message})
	}
}