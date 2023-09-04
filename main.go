package main

import (
	tgClient "bot/clients/telegram"
	event_consumer "bot/consumer/event-consumer"
	"bot/events/telegram"
	"bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize = 100
)

func main() {
	client := tgClient.New(tgBotHost, mustToken())
	eventsProcessor := telegram.New(
		&client,
		files.New(storagePath),
	)
	log.Print("Service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
