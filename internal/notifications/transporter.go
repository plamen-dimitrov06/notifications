package notifications

type Transporter interface {
	Send(message Message) bool
}
