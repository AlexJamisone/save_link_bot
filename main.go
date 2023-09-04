package main

import (
	tgClient "bot/clients/telegram"
	event_consumer "bot/consumer/event-consumer"
	"bot/events/telegram"
	"bot/storage/sqlite"
	"context"
	"flag"
	"log"
)

const (
	tgBotHost      = "api.telegram.org"
	storageSqlPath = "data/sqlite/storage.db"
	batchSize      = 100
)

func main() {
	client := tgClient.New(tgBotHost, mustToken())
	// s := files.New(storagePath)
	s, err := sqlite.New(storageSqlPath)
	if err != nil {
		log.Fatal("cant connect storage: ", err)
	}
	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("cant init storage", err)
	}
	eventsProcessor := telegram.New(
		&client,
		s,
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
