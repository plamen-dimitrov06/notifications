package notifications

import (
	"fmt"
	"time"
	"notifications/internal/transport"
)

type NotificationHandler struct {
	Transport transport.Transporter
}

func (h NotificationHandler) Notify(m transport.Message) {
	isSuccessful := h.Transport.Send(m)
	// if the message was not sent, try again in 3 seconds
	if isSuccessful != true {
		fmt.Printf("Unable to send message, retrying again.\n")
		time.Sleep(3 * time.Second)
		h.Notify(m)
	}
}

func NewEmailHandler() NotificationHandler {
	transport := transport.NewEmailTransport()
	return NotificationHandler{Transport: transport}
}

func NewSlackHandler() NotificationHandler {
	transport := transport.NewSlackTransport()
	return NotificationHandler{Transport: transport}
}

func NewSMSHandler() NotificationHandler {
	transport := transport.NewSMSTransport()
	return NotificationHandler{Transport: transport}
}