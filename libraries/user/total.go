package user

import (
	"context"
	"fmt"
	"log"

	"github.com/donghquinn/go_web/prisma/db"
)

func GetOneUsers(client *db.PrismaClient, email string) *db.ClientModel {
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

	return result
}