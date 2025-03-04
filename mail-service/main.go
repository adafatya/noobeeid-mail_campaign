package main

import (
	"log"
	"net/http"
)

// define port application
const (
	APP_PORT = ":8001"
)

func main() {
	// handle send mail
	http.HandleFunc("/send", sendEmail)

	log.Println("server running at port", APP_PORT)
	http.ListenAndServe(APP_PORT, nil)
}
