package main

import (
	"fmt"
	"github.com/pajlada/go-twitch-pubsub"
)

// const channelID = ""

func (cfg *apiConfig) pubsubRun(token string) {
	pubsubClient := twitchpubsub.NewClient(twitchpubsub.DefaultHost)
	pubsubClient.Listen(twitchpubsub.PointsEventTopic(cfg.ChannelID), token)

	pubsubClient.OnPointsEvent(pointsEventCallback)

	go pubsubClient.Start()
}

// Turn initiateElevenLabs into config method that blocks, and make a channel that will unblock when this is called
func pointsEventCallback(channelID string, data *twitchpubsub.PointsEvent) {
	if data.Reward.Title == "TTS" {
		fmt.Println("Got channel point message from topic!")
		initiateElevenLabs(data.UserInput)
	}
}
