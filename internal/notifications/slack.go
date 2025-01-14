package notifications

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackTransport struct {
}

func (t SlackTransport) Send(m Message) {
	api := slack.New("<API token here>")
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    m.Content,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	// TODO : move the hardcoded values to constant (channel and API key)
	channelID, timestamp, err := api.PostMessage(
		"<channel ID here>",
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}