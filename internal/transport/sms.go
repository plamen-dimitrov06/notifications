package transport

import (
	"context"
	"fmt"
	"os"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

type SMSTransport struct {
}

func (t SMSTransport) Send(m Message) bool {
	client, err := infobip.NewClient(os.Getenv("INFOBIP_URL"), os.Getenv("INFOBIP_API_TOKEN"))
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
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
		return false
	}
	return true
}

func NewSMSTransport() SMSTransport { return SMSTransport{} }