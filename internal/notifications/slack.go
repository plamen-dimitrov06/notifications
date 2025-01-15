package notifications

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type SlackTransport struct {
}

func (t SlackTransport) Send(m Message) {
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		m.Recipient,
		slack.MsgOptionText(m.Content, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func NewSlackTransport() SlackTransport { return SlackTransport{} }