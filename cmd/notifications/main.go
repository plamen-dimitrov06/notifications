package main

import (
	"net/http"
	"log"
	"notifications/internal/notifications"
)

func main() {
	http.HandleFunc("/slack", notifications.SlackHandler)
	http.HandleFunc("/sms", notifications.SMSHandler)
	http.HandleFunc("/email", notifications.EmailHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
