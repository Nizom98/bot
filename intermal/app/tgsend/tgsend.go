package tgsend

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New(bot *tgbotapi.BotAPI) *Send {
	return &Send{
		bot: bot,
	}
}

func (s *Send) Send(chatID int64, msg string) error {
	tgMsg := tgbotapi.NewMessage(chatID, msg)
	
	_, err := s.bot.Send(tgMsg)
	if err != nil {
		return fmt.Errorf("cannot send message: %w", err)
	}

	return nil
}

func (s *Send) SendTgMsg(msg tgbotapi.MessageConfig) error {	
	_, err := s.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("cannot send message: %w", err)
	}

	return nil
}