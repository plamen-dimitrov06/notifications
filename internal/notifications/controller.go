package notifications

import (
	"net/http"
	"encoding/json"
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

	// TODO: move to factory function
	transport := EmailTransport{}
	transport.Send(message)
}