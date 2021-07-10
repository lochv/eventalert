package main

import (
	winlog "github.com/bi-zone/gowinlog"
)

func inEvtIds(id uint64, evtIds []uint64) bool {
	for _, confId := range evtIds {
		if confId == id {
			return true
		}
	}
	return false
}

func main() {
	conf := loadConfig()
	teleBot := newTelegramBot(conf.TelegramBotToken, conf.TelegramChannelId)
	teleBot.run()
	teleBot.sendMsg("started")

	watcher, err := winlog.NewWinLogWatcher()
	if err != nil {
		teleBot.sendMsg(err.Error())
		return
	}

	if conf.EnableSysmon {
		err = watcher.SubscribeFromNow("Microsoft-Windows-Sysmon/Operational", "*")
		if err != nil {
			teleBot.sendMsg(err.Error())
			return
		}
	}

	err = watcher.SubscribeFromNow("Security", "*")
	if err != nil {
		teleBot.sendMsg(err.Error())
		return
	}

	for {
		select {
		case evt := <-watcher.Event():
			switch evt.Channel {
			case "Security":
				if inEvtIds(evt.EventId, conf.EvtIds) {
					teleBot.sendMsg(evt.Msg)
				}
			case "Microsoft-Windows-Sysmon/Operational":
				teleBot.sendMsg(evt.Msg)
			}
		}
	}
}
