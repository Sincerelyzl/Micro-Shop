package main

import (
	"context"
	"log"
	"microshop/config"
	"microshop/pkg/database"
	"microshop/server"
	"os"
)

func main() {
	ctx := context.Background()

	//Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error : .env path is required")
		}
		return os.Args[1]
	}())

	//DataBase connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	//Start server
	server.Start(ctx, &cfg, db)
}
