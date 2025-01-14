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

	// TODO : move the hardcoded values to constant (channel and API key)
	channelID, timestamp, err := api.PostMessage(
		m.Recipient,
		slack.MsgOptionText(m.Content, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}