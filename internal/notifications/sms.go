package notifications

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

type SMSTransport struct {
}

func (t SMSTransport) Send(m Message) {
	client, err := infobip.NewClient("4e1946.api.infobip.com", "3e8d5ed580b5664cf611e885a83283e2-8c6a15f2-d16d-4370-a721-58eaf7027fd8")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	sms := models.SMSMsg{
		Destinations: []models.SMSDestination{
			{To: m.Recipient},
		},
		From: "Notifications System",
		Text: m.Content,
	}
	request := models.SendSMSRequest{
		Messages: []models.SMSMsg{sms},
	}

	resp, respDetails, err := client.SMS.Send(context.Background(), request)
	fmt.Printf("Response : %s \r\n Response details : %s \r\n", resp, respDetails)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}