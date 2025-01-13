package notifications

import (
	"net/http"
	"encoding/json"
	"github.com/plamen-dimitrov06/notifications/internal/transport"
)

func SlackHandler(w http.ResponseWriter, r *http.Request) {
	message := NewMessage()
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SMSHandler(w http.ResponseWriter, r *http.Request) {
	message := NewMessage()
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	message := NewMessage()
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transport := transport.EmailTransport{}
	transport.Send(message)
}