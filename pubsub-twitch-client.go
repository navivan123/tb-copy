package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pajlada/go-twitch-pubsub"
	"log"
	"os"
)

const channelID = ""

func pubsubRun(token string) {
	godotenv.Load(".env")

	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("ELABS_API must be set")
	}

	pubsubClient := twitchpubsub.NewClient(twitchpubsub.DefaultHost)
	pubsubClient.Listen(twitchpubsub.PointsEventTopic(channelID), token)

	pubsubClient.OnPointsEvent(pointsEventCallback)

	go pubsubClient.Start()
}

func pointsEventCallback(channelID string, data *twitchpubsub.PointsEvent) {
	if data.Reward.Title == "TTS" {
		fmt.Println("Got channel point message from topic!")
		initiateElevenLabs(data.UserInput)
	}
}
