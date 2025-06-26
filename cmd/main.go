package main

import (
	"item-store/config"
	internal "item-store/internal/api"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	server, err := internal.NewServer(config)
	if err != nil {
		panic(err)
	}
	err = server.Start(":" + config.ApiPort)
	if err != nil {
		panic(err)
	}

}
