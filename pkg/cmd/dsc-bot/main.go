package dscbot

import (
	"gopkg.in/telebot.v3"

	logger "github.com/laerllik/dsc-bot/internal/logger"
	botcommands "github.com/laerllik/dsc-bot/pkg/bot-commands"
	tba "github.com/laerllik/dsc-bot/pkg/telegram-bot-api"
)

var COMMANDS = []botcommands.BotCommand{
	&botcommands.StartCommand{},
	&botcommands.CheckCommand{},
}

func Process() error {
	logger.Log("processing")

	commandsByChats := map[int64]botcommands.BotCommand{}

	tba.Bot.Self.Handle(telebot.OnText, func(c telebot.Context) error {
		chatId := c.Chat().ID

		command, exist := commandsByChats[chatId]

		if !exist {
			for _, commandPrototype := range COMMANDS {
				if commandPrototype.IsTriggered(c) {
					command = commandPrototype.Init()

					break
				}
			}
		}

		command.HandleMessage(c)

		if command.IsFinished() && exist {
			delete(commandsByChats, chatId)

			return nil
		}

		if !command.IsFinished() && !exist {
			commandsByChats[chatId] = command
		}

		return nil
	})

	tba.Bot.Self.Start()

	return nil
}
