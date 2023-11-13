package main

import (
	"net/http"

	"github.com/whatsauth/webhook"
)

func main() {
	http.HandleFunc("/", webhook.Post)
	http.ListenAndServe(":8080", nil)
}
