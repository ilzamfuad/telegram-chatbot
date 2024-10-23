package handlers

import (
	"telegram-chatbot/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	services.GetTodayCallback(bot, update)
}
