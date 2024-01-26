package data

import "github.com/donghquinn/go_web/prisma/db"

func GetPrisma() error {
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
    return err
  }

//     defer func() {
//     if err := client.Prisma.Disconnect(); err != nil {
//       panic(err)
//     }
//   }()

  return nil
}