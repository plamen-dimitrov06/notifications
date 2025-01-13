package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/slack", SlackHandler)
	http.HandleFunc("/sms", SMSHandler)
	http.HandleFunc("/email", EmailHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
