package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdHelp = "help"
	cmdList = "list"
)

type CmdRouter struct {
	bot *tgbotapi.BotAPI
	cmdHelp Cmder
	cmdList Cmder
}

type Cmder interface {
	Exec(mag *tgbotapi.Message) error
}

type Sender interface {
	Send(chatID int64, msg string) error
}
