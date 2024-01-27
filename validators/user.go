package validators

import (
	"github.com/donghquinn/go_web/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type requestFormValidationType struct {
	email string
	password string
}
var request types.UserLoginRequest

func UserGetOneValidator(ctx *gin.Context) types.UserLoginRequest {
	validator := validator.New()

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
}