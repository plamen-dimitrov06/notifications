package notifications

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailTransport struct {
}

func (t EmailTransport) Send(m Message) {
	c, err := smtp.Dial("mail:1025")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@notifications.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(m.Recipient); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	// TODO : missing subject, add it
	_, err = fmt.Fprintf(wc, m.Content)
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}

func NewEmailTransport() EmailTransport { return EmailTransport{} }