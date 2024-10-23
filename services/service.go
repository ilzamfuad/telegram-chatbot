package services

import (
	"telegram-chatbot/keyboards"
	"telegram-chatbot/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hi, disini bisa mengetahui cuaca di suatu kota."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func GetToday(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Masukan kota tempat yang ingin tahu cuacanya."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func GetTodayCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	infoWeather, err := repositories.GetWeather(update)
	if err != nil {
		infoWeather = "Tidak bisa mendapatkan info cuaca"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, infoWeather)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
