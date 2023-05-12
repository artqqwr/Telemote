package main

import (
	"github.com/artqqwr/telemote/internal/bot"
	"github.com/artqqwr/telemote/internal/config"
	"github.com/artqqwr/telemote/internal/database"
	"github.com/artqqwr/telemote/internal/repository"
	"github.com/artqqwr/telemote/internal/server"
)

func main() {
	config := config.Load()
	database := database.New(config.Database)
	db := database.Connect()

	_ = repository.New(db)

	bot := bot.New(
		bot.WithConfig(config.Bot),
	)
	server := server.New(config.Server)

	go bot.Run()
	server.Run()
}
