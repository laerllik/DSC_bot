package botcommands

import (
	"gopkg.in/telebot.v3"
)

type BotCommand interface {
	IsTriggered(c telebot.Context) bool
	Init() BotCommand
	IsFinished() bool
	HandleMessage(c telebot.Context) error
}
