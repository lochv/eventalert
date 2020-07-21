package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)
type TelegramBot struct {
	Bot *tgbotapi.BotAPI
	ChannelId string
}
func (t *TelegramBot) NewBot(Token string) {
	t.Bot, _ = tgbotapi.NewBotAPI(Token)
}

func (t *TelegramBot) SendReport(Data string) {
	msg := tgbotapi.NewMessageToChannel(t.ChannelId, Data)
	msg.DisableWebPagePreview = true
	_, err := t.Bot.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}