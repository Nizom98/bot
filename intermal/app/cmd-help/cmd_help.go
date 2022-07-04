package cmdhelp

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func New(send Sender) *Help {
	return &Help{
		send: send,
	}
}

func (cmd *Help) Exec(msg *tgbotapi.Message) error {
	return cmd.send.Send(msg.Chat.ID, "no info for help")
}