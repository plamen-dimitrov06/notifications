package transport

import (
	"fmt"
	"net/smtp"
)

type EmailTransport struct {
}

func (t EmailTransport) Send(m Message) bool {
	c, err := smtp.Dial("mail:1025")
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@notifications.com"); err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	if err := c.Rcpt(m.Recipient); err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	wc, err := c.Data()
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	_, err = fmt.Fprintf(wc, m.Content)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	err = wc.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	return true
}

func NewEmailTransport() EmailTransport { return EmailTransport{} }