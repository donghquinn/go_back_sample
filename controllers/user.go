package controllers

import (
	"log"

	"github.com/donghquinn/go_web/data"
	"github.com/donghquinn/go_web/dto"
	"github.com/donghquinn/go_web/libraries/user"
	"github.com/gin-gonic/gin"
)

func UserCountControllers(ctx *gin.Context){

	dbClient, prismaErr := data.GetPrisma()

	if prismaErr != nil {
		log.Fatalln(prismaErr)
	}

	count := user.GetTotalUsers(dbClient)

	dto.SetResponse(200, count, ctx)
}