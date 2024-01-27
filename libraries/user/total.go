package user

import (
	"context"
	"fmt"
	"log"

	"github.com/donghquinn/go_web/prisma/db"
)

type QueriedUserInfo struct {
	Name string `json:"name" xml:"name" binding:"required"`
	Email string `json:"email" xml:"email" binding:"required"`
}

func GetOneUsers(client *db.PrismaClient, email string) QueriedUserInfo {
	ctx := context.Background()

	result, queryErr := client.Client.FindFirst(
		db.Client.Email.Equals(email),
	).Exec(ctx)

	if queryErr != nil {
		log.Fatalln(queryErr)
	}
	
	defer func() {
   	 if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()

	fmt.Printf("Result: %d", len(result.Name))

	queried := QueriedUserInfo {
		Name: result.Name,
		Email: result.Email,
	}

	return queried
}