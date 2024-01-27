package dto

import (
	"github.com/donghquinn/go_web/types"
	"github.com/gin-gonic/gin"
)

// type KeyableObject struct {
// 	data any
// }

func SetResponse(resCode int, data any, ctx *gin.Context) {
	// response := ResponseType {
	// 	resCode: resCode,
	// 	dataRes:data,
	// 	errMsg: make([]string, 0),
	// }
	// dataResResult := DataResResult {
	// 	result: data,
	// }

	ctx.JSON(resCode, types.ResponseType {
		ResCode: resCode,
		DataRes: data,
	})

	return 
}