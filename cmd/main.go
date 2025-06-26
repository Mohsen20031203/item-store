package main

import (
	"item-store/config"
	_ "item-store/docs"
	internal "item-store/internal/api"
	"log"
)

// @title           My API
// @version         1.0
// @description     This is a sample API for managing items.
// @host            localhost:1010
// @BasePath        /
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
