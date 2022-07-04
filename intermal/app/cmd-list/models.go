package cmdlist

import "github.com/Nizom98/bot/intermal/service/product"

type Sender interface {
	Send(chatID int64, msg string) error
}

type List struct {
	products []*product.Product
	send     Sender
}