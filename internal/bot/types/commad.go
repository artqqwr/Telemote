package types

import (
	tele "gopkg.in/telebot.v3"
)

type Command struct {
	Name       string
	About      string
	HandleFunc tele.HandlerFunc
}

type Commands []*Command
