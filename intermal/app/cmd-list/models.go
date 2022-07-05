package cmdlist

import (
	"github.com/Nizom98/bot/intermal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender interface {
	Send(chatID int64, msg string) error
	SendTgMsg(msg tgbotapi.MessageConfig) error
}

type List struct {
	products []*product.Product
	send     Sender
}