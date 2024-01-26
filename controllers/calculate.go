package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/go_web/libraries/calculator"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	Num1 int64 `json:"num1"`
	Num2 int64 `json:"num2"`
}

func CalculateAddController(ctx *gin.Context) {
	request := CalculateRequest{}

	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"errorMessage": "[ADD] Required Body Not Included"})
	} else {
		num1Value, numErr1 := strconv.ParseInt(ctx.PostForm("num1"), 10, 8)
		if numErr1 != nil {
			log.Fatalln(numErr1)
		}

		fmt.Printf("Parsed Num1 Value: %d", num1Value)

		num2Value, numErr2 := strconv.ParseInt(ctx.PostForm("num2"), 10, 8)

		if numErr2 != nil {
			log.Fatalln(numErr2)
		}
		fmt.Printf("Parsed Num1 Value: %d", num1Value)

		value := calculator.Add(float64(num1Value), float64(num2Value))

		ctx.JSON(http.StatusOK, gin.H{"Added Value": value})
	}
}

func CalculateSumController(ctx *gin.Context) {
		request := CalculateRequest{}

	if err := ctx.ShouldBind(request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"errorMessage": "[SUM] Required Body Not Included"})
	} else {
		num1Value, numErr1 := strconv.ParseInt(ctx.PostForm("num1"), 10, 8)
		if numErr1 != nil {
			log.Fatalln(numErr1)
		}

		fmt.Printf("[SUM] Parsed Num1 Value: %d", num1Value)

		num2Value, numErr2 := strconv.ParseInt(ctx.PostForm("num2"), 10, 8)

		if numErr2 != nil {
			log.Fatalln(numErr2)
		}
		fmt.Printf("[SUM] Parsed Num1 Value: %d", num1Value)

		value, sumErr := calculator.Sum(int(num1Value), int(num2Value))

		if sumErr != nil {
			log.Fatalln(sumErr)
		}

		ctx.JSON(http.StatusOK, gin.H{"Added Value": value})
	}
}