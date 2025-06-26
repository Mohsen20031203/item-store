package main

import (
	"item-store/config"
	internal "item-store/internal/api"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server, err := internal.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(":" + config.ApiPort)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	select {}

}
