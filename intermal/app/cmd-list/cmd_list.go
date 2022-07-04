package cmdlist

import (
	"fmt"
	"strings"

	"github.com/Nizom98/bot/intermal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New(send Sender, products []*product.Product) *List {
	return &List{
		send: send,
		products: products,
	}
}

func (cmd *List) Exec(msg *tgbotapi.Message) error {
	return cmd.send.Send(msg.Chat.ID,productsToString(cmd.products))
}

func productsToString(products []*product.Product) string {
	builder := strings.Builder{}
	
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%s\n", p.Title))
	}

	return builder.String()
}