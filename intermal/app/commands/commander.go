package commands

import (
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
	}
}

func (cmd *CmdRouter) HandleUpdate(msg *tgbotapi.Message) error {
	if msg == nil {
		return nil
	}
	switch msg.Command() {
	case cmdHelp: 
		return cmd.cmdHelp.Exec(msg)	
	case cmdList: 
		return cmd.cmdList.Exec(msg)
	default:
		return nil
	}
}