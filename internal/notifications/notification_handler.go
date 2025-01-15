package notifications

import (
	"fmt"
	"time"
)

type NotificationHandler struct {
	Transport Transporter
}

func (h NotificationHandler) Notify(m Message) {
	isSuccessful := h.Transport.Send(m)
	// if the message was not sent, try again in 5 seconds
	if isSuccessful != true {
		fmt.Printf("Unable to send message, retrying again.\n")
		time.Sleep(5 * time.Second)
		h.Notify(m)
	}
}

func NewEmailHandler() NotificationHandler {
	transport := NewEmailTransport()
	return NotificationHandler{Transport: transport}
}

func NewSlackHandler() NotificationHandler {
	transport := NewSlackTransport()
	return NotificationHandler{Transport: transport}
}

func NewSMSHandler() NotificationHandler {
	transport := NewSMSTransport()
	return NotificationHandler{Transport: transport}
}