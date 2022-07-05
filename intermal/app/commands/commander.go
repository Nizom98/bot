package commands

import (
	"log"

	cmdget "github.com/Nizom98/bot/intermal/app/cmd-get"
	cmdhelp "github.com/Nizom98/bot/intermal/app/cmd-help"
	cmdlist "github.com/Nizom98/bot/intermal/app/cmd-list"
	"github.com/Nizom98/bot/intermal/app/tgsend"
	"github.com/Nizom98/bot/intermal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New(bot *tgbotapi.BotAPI) *CmdRouter {
	send := tgsend.New(bot)
	products := product.List()

	return &CmdRouter{
		bot: bot,
		cmdHelp: cmdhelp.New(send),
		cmdList: cmdlist.New(send, products),
		cmdGet: cmdget.New(send),
	}
}

func (cmd *CmdRouter) HandleUpdate(upd tgbotapi.Update) error {
	if upd.Message == nil {
		return nil
	}
	log.Println("got update")

	if upd.CallbackQuery != nil {
		callbackData := upd.CallbackQuery.Data
		log.Println("got collback data: ", callbackData)
	}
	switch upd.Message.Command() {
	case cmdHelp: 
		return cmd.cmdHelp.Exec(upd.Message)	
	case cmdList: 
		return cmd.cmdList.Exec(upd.Message)
	case cmdGet: 
		return cmd.cmdGet.Exec(upd.Message)
	default:
		return nil
	}
}