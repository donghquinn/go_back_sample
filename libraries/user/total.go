package user

import (
	"context"
	"fmt"
	"log"

	"github.com/donghquinn/go_web/prisma/db"
)

func GetTotalUsers(client *db.PrismaClient) []db.ClientModel {
	ctx := context.Background()
	result, queryErr := client.Client.FindMany().Exec(ctx)

	if queryErr != nil {
		log.Fatalln(queryErr)
	}
	
	defer func() {
   	 if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()

	fmt.Printf("Result: %d", len(result))

	return result
}