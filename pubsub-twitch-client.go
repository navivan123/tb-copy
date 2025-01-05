package main

import (
	"fmt"
	"github.com/pajlada/go-twitch-pubsub"
	"math/rand"
	"strconv"
)

// Prepares the PubSub client to be used in other functions
func (cfg *apiConfig) initTwitchPubSub() {

	pubsubClient := twitchpubsub.NewClient(twitchpubsub.DefaultHost)
	cfg.TwitchPubSubClient = pubsubClient

}

// Functions to configure the client to receive various different pubsub events
func (cfg *apiConfig) twitchPubSubListenToPointsEvents() {

	pubsubClient := cfg.TwitchPubSubClient

	pubsubClient.Listen(twitchpubsub.PointsEventTopic(cfg.ChannelID), cfg.TwitchToken)
	pubsubClient.OnPointsEvent(cfg.pointsEventCallback)

}

func (cfg *apiConfig) twitchPubSubListenToModEvents() {

	pubsubClient := cfg.TwitchPubSubClient

	pubsubClient.Listen(twitchpubsub.ModerationActionTopic(cfg.ChannelID, cfg.ChannelID), cfg.TwitchToken)
	pubsubClient.OnModerationAction(cfg.modEventCallback)

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

func (cfg *apiConfig) modEventCallback(channelID string, data *twitchpubsub.ModerationAction) {
	if data.ModerationAction == "timeout" {
		a := map[int]string{0: "ooh.mp3", 1: "oohh.mp3"}
		playAudio(a[rand.Intn(len(a))])
		return
	}

	path := ""
	if data.ModerationAction == "" {
		path = path + "mod"
	} else {
		path = data.ModerationAction
	}

	path = path + ".mp3"
	playAudio(path)

}
