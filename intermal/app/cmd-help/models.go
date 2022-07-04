package cmdhelp

type Sender interface {
	Send(chatID int64, msg string) error
}

type Help struct {
	send Sender
}