package botcommands

import (
	"gopkg.in/telebot.v3"
)

type StartCommand struct{}

func (cmd *StartCommand) IsTriggered(c telebot.Context) bool {
	return c.Message().Text == "/start"
}

func (cmd *StartCommand) Init() BotCommand {
	return &StartCommand{}
}

func (cmd *StartCommand) IsFinished() bool {
	return true
}

func (cmd *StartCommand) HandleMessage(c telebot.Context) error {
	_, err := c.Bot().Send(c.Chat(), "Hello")

	return err
}
