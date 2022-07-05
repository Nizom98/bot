package cmdget

type Sender interface {
	Send(chatID int64, msg string) error
}

type Get struct {
	send Sender
}