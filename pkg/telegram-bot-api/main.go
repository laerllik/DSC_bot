package telegrambotapi

import (
	"time"

	telebot "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	Self *telebot.Bot
}

func (tb *TelegramBot) Configure(token string) error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return err
	}

	tb.Self = bot

	return nil
}

var Bot = TelegramBot{}
