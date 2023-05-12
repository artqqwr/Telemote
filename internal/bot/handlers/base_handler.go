package handlers

import (
	"fmt"

	"github.com/artqqwr/telemote/internal/bot/types"
	tele "gopkg.in/telebot.v3"
)

type BaseHandler struct {
	commands *types.Commands
}

func NewBaseHandler(commands *types.Commands) *BaseHandler {
	return &BaseHandler{
		commands: commands,
	}
}

func (h *BaseHandler) Help(c tele.Context) error {
	message := ""

	for _, command := range *h.commands {
		message += fmt.Sprintf("%s — %s\n", command.Name, command.About)
	}

	return c.Send(message)
}

func (h *BaseHandler) Start(c tele.Context) error {
	return c.Send("Hello qweqwe!!!")
}
