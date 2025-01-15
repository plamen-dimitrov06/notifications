package notifications

type NotificationHandler struct {
	Transport Transporter
}

func (h NotificationHandler) Notify(m Message) {
	h.Transport.Send(m)
	// err := h.Transport.Send(m)
	// // if the message was not sent, try again
	// if err != nil {
	// 	h.Notify(m)
	// }
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