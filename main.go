package main

import (
	"fmt"
	"github.com/haguro/elevenlabs-go"
	"github.com/joho/godotenv"
	"github.com/pajlada/go-twitch-pubsub"
	"log"
	"os"
)

type apiConfig struct {
	TwitchToken        string
	ChannelID          string
	ClientID           string
	ElabsKey           string
	ElabsClient        *elevenlabs.Client
	TwitchPubSubClient *twitchpubsub.Client
}

func main() {

	// ELABS API KEY for Elabs API requests
	godotenv.Load(".env")
	elevenAPI := os.Getenv("ELABS_API")
	if elevenAPI == "" {
		log.Fatal("ELABS_API must be set")
	}

	// Getting channel ID to pull PubSub info from
	godotenv.Load(".env")
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("CHANNEL_ID must be set")
	}

	// Getting Client_ID as all Twitch API calls need this set
	godotenv.Load(".env")
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("CLIENT_ID must be set")
	}

	config := apiConfig{TwitchToken: "", ClientID: clientID, ChannelID: channelID, ElabsKey: elevenAPI, ElabsClient: nil, TwitchPubSubClient: nil}

	// Initializes all the services the bot uses (twitch api, elabs, pubsub...)
	config.initEleven()
	config.getAuthToken()
	config.initTwitchPubSub()

	// Start all "daemons"
	config.twitchPubSubListenToPointsEvents()

	// Wait for user input to exit
	fmt.Println("Press any key or Ctrl+C to stop!")
	fmt.Scanln()
}
