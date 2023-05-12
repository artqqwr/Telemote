package bot

import (
	"github.com/artqqwr/telemote/internal/bot/handlers"
	"github.com/artqqwr/telemote/internal/bot/types"
	"github.com/artqqwr/telemote/internal/repository"
	tele "gopkg.in/telebot.v3"
)

type Option func(*Bot)
type Bot struct {
	api      *tele.Bot
	config   *Config
	commands *types.Commands
}

func WithRepository(repo repository.Repository) Option {
	return func(b *Bot) {

	}
}

func WithConfig(config *Config) Option {
	return func(b *Bot) {
		b.config = config
	}
}

func New(opts ...Option) *Bot {

	bot := &Bot{
		commands: &types.Commands{},
	}

	for _, opt := range opts {
		opt(bot)
	}

	api, err := tele.NewBot(bot.config.Settings)
	if err != nil {
		panic(err)
	}
	bot.api = api

	// Command handlers
	baseHandler := handlers.NewBaseHandler(bot.commands)

	initRoutes(bot, baseHandler)

	return bot
}

func (b *Bot) Run() {
	b.api.Start()
}

func (b *Bot) AddCommand(name string, handleFunc tele.HandlerFunc, about string) {
	*b.commands = append(*b.commands, &types.Command{
		Name:       name,
		About:      about,
		HandleFunc: handleFunc,
	})
}
