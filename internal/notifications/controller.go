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

	handler := NewSlackHandler()
	handler.Notify(message)
}

func SMSHandler(w http.ResponseWriter, r *http.Request) {
	message := NewMessage()
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler := NewSMSHandler()
	handler.Notify(message)
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	message := NewMessage()
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	handler := NewEmailHandler()
	handler.Notify(message)
}