package main

import (
	"bufalo-wallet-app/config"
	"bufalo-wallet-app/database"
	"bufalo-wallet-app/routes"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDatabase(cfg)

	router := routes.SetupRouter()
	log.Fatal(router.Run(cfg.ServerPort))
}
