package dto

import (
	"github.com/gin-gonic/gin"
)

// type KeyableObject struct {
// 	data any
// }

type DataResResult struct {
	result any
}

type ResponseType  struct {
	resCode int
	dataRes DataResResult
	errMsg []string
}

func SetResponse(resCode int, data any, ctx *gin.Context) {
	// response := ResponseType {
	// 	resCode: resCode,
	// 	dataRes:data,
	// 	errMsg: make([]string, 0),
	// }
	// dataResResult := DataResResult {
	// 	result: data,
	// }

	ctx.JSON(resCode, gin.H{	
		"resCode": resCode,
		"dataRes": data,
		"errMsg": make([]string, 0),})

	return 
}