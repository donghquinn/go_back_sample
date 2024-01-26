package data

import "github.com/donghquinn/go_web/prisma/db"

func GetPrisma() (*db.PrismaClient,error ){
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
    return nil, err
  }

  return client, nil
}