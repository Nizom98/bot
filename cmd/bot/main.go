package main

import (
	"log"
	"os"

	"github.com/Nizom98/bot/intermal/app/commands"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	godotenv.Load()

	token := os.Getenv("token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	cmd := commands.New(bot)

	for update := range updates {
		cmd.HandleUpdate(update.Message)
	}
}

// time 02:25:43