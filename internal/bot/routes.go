package bot

import "github.com/artqqwr/telemote/internal/bot/handlers"

func initRoutes(bot *Bot, baseHandler *handlers.BaseHandler) {
	bot.AddCommand("/start", baseHandler.Start, "начать")
	bot.AddCommand("/help", baseHandler.Help, "получить список команд:\n")

	for _, command := range *bot.commands {
		bot.api.Handle(command.Name, command.HandleFunc)
	}
}
