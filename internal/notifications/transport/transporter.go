package transport

type Transporter interface {
	Send(message Message)
}
