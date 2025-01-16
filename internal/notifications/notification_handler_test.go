package notifications

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"notifications/internal/transport"
)

func TestFactoryFunctionsWorkCorrectly(t *testing.T) {
	tests := map[string]struct {
		input  NotificationHandler
		expected transport.Transporter
	}{
		"Validate the email handler has the email transport.": {
			input: NewEmailHandler(),
			expected: transport.NewEmailTransport(),
		},
		"Validate the slack handler has the slack transport.": {
			input: NewSlackHandler(),
			expected: transport.NewSlackTransport(),
		},
		"Validate the SMS handler has the SMS transport.": {
			input: NewSMSHandler(),
			expected: transport.NewSMSTransport(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.input.Transport, test.expected)
		})
	}
}
