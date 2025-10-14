package main

import (
	"os"

	logger "github.com/laerllik/dsc-bot/internal/logger"
	tba "github.com/laerllik/dsc-bot/pkg/telegram-bot-api"

	dscbot "github.com/laerllik/dsc-bot/pkg/cmd/dsc-bot"
)

func main() {
	logger.Configure("dsc-bot. ")
	logger.Log("start initialize")
	tba.Bot.Configure(os.Getenv("TOKEN"))

	logger.Log("finish initialize")
	logger.Log("start process")
	dscbot.Process()
	logger.Log("finish process")
}
