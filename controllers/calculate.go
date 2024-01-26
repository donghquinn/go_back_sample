package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/donghquinn/go_web/libraries/calculator"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	Num1 int `form:"num1" json:"num1" xml:"num1" binding:"required"`
	Num2 int `form:"num2" json:"num2" xml:"num2" binding:"required"`
}

func CalculateAddController(ctx *gin.Context) {
	var request CalculateRequest

	if bodyErr := ctx.ShouldBind(&request); bodyErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is Not Included","error": bodyErr.Error()})

		return
	}  
	num1 := request.Num1
	num2 := request.Num2

	fmt.Printf("[ADD] Body Data Num1: %o, Num2: %o\n", num1, num2)


	value := calculator.Add(float64(num1), float64(num2))

	ctx.JSON(http.StatusOK, gin.H{"Added Value": value})
}


func CalculateSumController(ctx *gin.Context) {
	var request CalculateRequest

	if bodyErr := ctx.ShouldBind(&request); bodyErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is Not Included","error": bodyErr.Error()})

		return
	} 

	num1 := request.Num1
	num2 := request.Num2
	
	fmt.Printf("[SUM] Body Data Num1: %o, Num2: %o\n", num1, num2)

	value, sumErr := calculator.Sum(num1, num2)

	if sumErr != nil {
		log.Fatalln(sumErr)
	}

	ctx.JSON(http.StatusOK, gin.H{"Added Value": value})
	
	
	
}