package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	TelegramChannelId string   `yaml:"telegram-channel-id"`
	TelegramBotToken  string   `yaml:"telegram-bot-token"`
	EvtIds            []uint64 `yaml:"evt-ids"`
	EnableSysmon      bool     `yaml:"enable-sysmon"`
}

func loadConfig() config {
	yamlFile, err := ioutil.ReadFile("C:\\Windows\\evt_alert.yaml")
	if err != nil {
		panic("Missing C:\\Windows\\evt_alert.yaml")
	}
	var evtAlertConfig config
	err = yaml.Unmarshal(yamlFile, &evtAlertConfig)
	if err != nil {
		panic(err.Error())
	}
	return evtAlertConfig
}
