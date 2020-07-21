package main

import (
	"fmt"
	winlog "github.com/scalingdata/gowinlog"
)

func main() {
	watcher, err := winlog.NewWinLogWatcher()
	if err != nil {
		fmt.Printf("Couldn't create watcher: %v\n", err)
		return
	}

	watcher.SubscribeFromNow("Microsoft-Windows-Sysmon/Operational","*")
	watcher.SubscribeFromNow("Security","*")
	
	bot := TelegramBot{
		Bot:       nil,
		// create channel => get channel id
		ChannelId: "",
	}
	// create new bot with botfather => add bot to channel 's admin => get bot 's token
	bot.NewBot("")
	for {
		select {
		case evt := <- watcher.Event():
			switch evt.Channel {
			case "Security":
				switch evt.EventId {
				case 4648:
					//log on success
					bot.SendReport(evt.Msg)
				case 4625:
					//log on fail
					bot.SendReport(evt.Msg)
				case 4663:
					//file access
					bot.SendReport(evt.Msg)
				case 5145:
					//file access
					bot.SendReport(evt.Msg)
				}
			case "Microsoft-Windows-Sysmon/Operational":
				bot.SendReport(evt.Msg)
			}
		}
	}
}