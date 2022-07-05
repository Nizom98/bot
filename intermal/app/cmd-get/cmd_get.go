package cmdget

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New(send Sender) *Get {
	return &Get{
		send: send,
	}
}

func (cmd *Get) Exec(msg *tgbotapi.Message) error {
	args := msg.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("cannot convert 'get' command argument: ", err)
		return cmd.send.Send(msg.Chat.ID, "cannot convert command argument")
	}

	
	return cmd.send.Send(msg.Chat.ID, fmt.Sprintf("parsed argument: %d", arg))
}