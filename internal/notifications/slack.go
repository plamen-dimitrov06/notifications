package notifications

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type SlackTransport struct {
}

func (t SlackTransport) Send(m Message) bool {
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		m.Recipient,
		slack.MsgOptionText(m.Content, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true
}

func NewSlackTransport() SlackTransport { return SlackTransport{} }