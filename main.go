package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type apiConfig struct {
	TwitchToken string
	ChannelID   string
	ClientID    string
	ElabsKey    string
	ElabsClient *elevenlabs.Client
}

func main() {

	godotenv.Load(".env")

	elevenAPI := os.Getenv("ELABS_API")
	if elevenAPI == "" {
		log.Fatal("ELABS_API must be set")
	}

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

	apiConfigForManyDifferentAPIsThatThisApplicationUsesForAuthenticatingAndPubSubAndStuff := apiConfig{TwitchToken: "", ClientID: clientID, ChannelID: channelID, ElabsKey: elevenAPI}

	apiConfigForManyDifferentAPIsThatThisApplicationUsesForAuthenticatingAndPubSubAndStuff.getAuthToken()
	apiConfigForManyDifferentAPIsThatThisApplicationUsesForAuthenticatingAndPubSubAndStuff.pubsubRun()
	fmt.Println("Press any key or Ctrl+C to stop!")
	fmt.Scanln()
}
