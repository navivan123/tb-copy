package main

import (
	"fmt"
	"github.com/pajlada/go-twitch-pubsub"
	"strconv"
)

// Prepares the PubSub client
func (cfg *apiConfig) initTwitchPubSub() {
	pubsubClient := twitchpubsub.NewClient(twitchpubsub.DefaultHost)
	cfg.TwitchPubSubClient = pubsubClient

}

// Configures the client to receive twitch point events
func (cfg *apiConfig) twitchPubSubListenToPointsEvents() {

	pubsubClient := cfg.TwitchPubSubClient

	pubsubClient.Listen(twitchpubsub.PointsEventTopic(cfg.ChannelID), cfg.TwitchToken)
	pubsubClient.OnPointsEvent(cfg.pointsEventCallback)

	go pubsubClient.Start()
}

// Parses and plays the message received by the TTS Points Channel Reward
func (cfg *apiConfig) pointsEventCallback(channelID string, data *twitchpubsub.PointsEvent) {
	if data.Reward.Title == "TTS" {
		fmt.Println("Got channel point message from topic!")
		arr := replace(data.UserInput)
		for _, voice := range arr {
			idx, _ := strconv.Atoi(voice[0:1])
			callEleven(cfg.ElabsClient, getVoices()[idx], voice[1:])
		}
	}
}
