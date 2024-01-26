package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/go_web/data"
	"github.com/donghquinn/go_web/dto"
	"github.com/donghquinn/go_web/libraries/user"
	"github.com/donghquinn/go_web/types"
	"github.com/gin-gonic/gin"
)

func UserCountControllers(ctx *gin.Context){
	request := types.UserLoginRequest{}

	if reqErr:=ctx.ShouldBind(&request); reqErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is Not Included", "error": reqErr.Error()})

		return
	}

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	log.Printf("Email: %s, Password: %s\n", email, password)
	dbClient, prismaErr := data.GetPrisma()

	if prismaErr != nil {
		log.Fatalln(prismaErr)
	}

	count := user.GetTotalUsers(dbClient)

	dto.SetResponse(200, count, ctx)
}