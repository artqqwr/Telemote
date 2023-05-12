package config

import (
	"time"

	"github.com/artqqwr/telemote/internal/bot"
	"github.com/artqqwr/telemote/internal/database"
	"github.com/artqqwr/telemote/internal/server"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

type Config struct {
	Bot      *bot.Config
	Server   *server.Config
	Database *database.Config
}

func Load() *Config {
	return &Config{
		Bot: &bot.Config{
			Settings: telebot.Settings{
				Token:  "token",
				Poller: &tele.LongPoller{Timeout: 10 * time.Second},
				OnError: func(error, telebot.Context) {
				},
			},
		},
		Server: &server.Config{
			Port: 3000,
		},
		Database: &database.Config{
			Dsn:    "./database.db",
			Driver: "sqlite",
		},
	}
}
