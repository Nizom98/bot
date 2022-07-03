package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Nizom98/bot/intermal/service/product"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdHelp = "help"
	cmdList = "list"
)

var productsSrv = product.NewService()

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

	for update := range updates {
		handleMessage(bot, update.Message)
	}
}


func send(bot *tgbotapi.BotAPI, chatID int64, msg string) error {
	tgMsg := tgbotapi.NewMessage(chatID, msg)
	_, err := bot.Send(tgMsg)
	return err
}

func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	if msg == nil {
		return nil
	}
	switch msg.Command() {
	case cmdHelp: 
		return send(bot, msg.Chat.ID, "no info for help")	
	case cmdList: 
		return send(bot, msg.Chat.ID, productsToString(productsSrv.List()))	
	default:
		return send(bot, msg.Chat.ID, msg.Text)
	}
}

func productsToString(products []*product.Product) string {
	builder := strings.Builder{}
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%s\n", p.Title))
	}

	return builder.String()
}
// time 01:52:43