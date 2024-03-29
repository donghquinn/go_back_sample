package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/donghquinn/go_web/dto"
	"github.com/donghquinn/go_web/libraries/calculator"
	"github.com/donghquinn/go_web/types"
	"github.com/gin-gonic/gin"
)

// Request Type Definition


func CalculateAddController(ctx *gin.Context) {
	var request types.CalculateRequest

	if bodyErr := ctx.ShouldBind(&request); bodyErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is Not Included", "error": bodyErr.Error()})

		return
	}

	num1 := request.Num1
	num2 := request.Num2

	fmt.Printf("[ADD] Body Data Num1: %o, Num2: %o\n", num1, num2)

	value := calculator.Add(float64(num1), float64(num2))


	dto.SetResponse(200, value, ctx)
	// ctx.JSON(http.StatusOK, gin.H{"addValue": value})
}


func CalculateSumController(ctx *gin.Context) {
	var request types.CalculateRequest

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
	dto.SetResponse(200, value, ctx)
}