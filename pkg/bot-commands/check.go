package botcommands

import (
	"gopkg.in/telebot.v3"
)

type CheckCommand struct {
	stage string
}

func (cmd *CheckCommand) IsTriggered(c telebot.Context) bool {
	return c.Message().Text == "/check"
}

func (cmd *CheckCommand) Init() BotCommand {
	return &CheckCommand{stage: "greet"}
}

func (cmd *CheckCommand) IsFinished() bool {
	return cmd.stage == "finish"
}

func (cmd *CheckCommand) HandleMessage(c telebot.Context) error {
	if cmd.stage == "greet" {
		cmd.stage = "ask-about-himself"
		_, err := c.Bot().Send(c.Chat(), "Hello! How are you?")

		return err
	}

	if cmd.stage == "ask-about-himself" {
		cmd.stage = "say-goodbye"
		_, err := c.Bot().Send(c.Chat(), "Nice to hear. What are you doing?")

		return err
	}

	if cmd.stage == "say-goodbye" {
		cmd.stage = "finish"
		_, err := c.Bot().Send(c.Chat(), "Bye")

		return err
	}

	return nil
}
