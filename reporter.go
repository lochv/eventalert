package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegramBot struct {
	bot       *tgbotapi.BotAPI
	channelId string
	msgChan   chan string
}

func newTelegramBot(token string, channelId string) *telegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err.Error())
	}
	return &telegramBot{
		bot:       bot,
		channelId: channelId,
		msgChan:   make(chan string, 16),
	}
}

func (t *telegramBot) run() {
	go func() {
		for {
			mess := <-t.msgChan
			msg := tgbotapi.NewMessageToChannel(t.channelId, mess)
			msg.DisableWebPagePreview = true
			_, err := t.bot.Send(msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}

func (t *telegramBot) sendMsg(mess string) {
	t.msgChan <- mess
}
