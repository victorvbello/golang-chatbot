package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorvbello/golang-chatbot/bot"
	"github.com/victorvbello/golang-chatbot/chat"
)

const WEB_SERVER_PORT = "3001"

func main() {
	chatbot := bot.NewAgentCase()
	hub := chat.NewHub(chatbot)

	go hub.Run()

	r := mux.NewRouter()

	r.Handle("/ws", chat.ServeWs(hub))

	http.Handle("/", r)

	log.Print("Init")

	http.ListenAndServe(":"+WEB_SERVER_PORT, nil)
}
