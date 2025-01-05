package main

import (
	"context"
	//"fmt"
	"github.com/haguro/elevenlabs-go"
	"log"
	"os"
	"os/exec"
	"time"
)

const modelID = "eleven_multilingual_v2"

// Initiates elevenlabs client and returns it to use later
func (cfg *apiConfig) initEleven() {
	client := elevenlabs.NewClient(context.Background(), cfg.ElabsKey, 15*time.Second)
	cfg.ElabsClient = client
	return
}

// Calls elevenlabs API to generate mp3 with text and voice, and then play it
func callEleven(client *elevenlabs.Client, voice, text string) {
	log.Printf("Making Request Struct\n")
	ttsReq := elevenlabs.TextToSpeechRequest{
		Text:    text,
		ModelID: modelID}

	log.Printf("Calling client\n")
	audio, err := client.TextToSpeech(voice, ttsReq)
	if err != nil {
		log.Printf("ERROR CALLING CLIENT\n")
		log.Fatal(err)
	}

	writeAudioAndPlay(audio)

}

func writeAudioAndPlay(audio []byte) {
	log.Printf("Writing File\n")
	file, err := os.Create("elabs.mp3")
	if err != nil {
		log.Printf("ERROR CREATING FILE\n")
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write(audio)
	if err != nil {
		log.Printf("ERROR WRITING TO FILE\n")
		log.Fatal(err)
	}
	err = file.Sync()
	if err != nil {
		log.Printf("ERROR WRITING TO FILE\n")
		log.Fatal(err)
	}

	log.Printf("Executing File\n")
	cmd := exec.Command("mpg123", "elabs.mp3")
	if err = cmd.Start(); err != nil {
		log.Printf("ERROR EXECUTING FILE\n")
		log.Fatal(err)
	}
	if err = cmd.Wait(); err != nil {
		log.Printf("ERROR WAITING FOR COMMAND FILE")
		log.Fatal(err)
	}
}
