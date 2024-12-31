package main

import (
	"context"
	//"fmt"
	"github.com/haguro/elevenlabs-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"time"
)

const modelID = "eleven_multilingual_v2"

func initEleven() *elevenlabs.Client {
	godotenv.Load(".env")

	elevenAPI := os.Getenv("ELABS_API")
	if elevenAPI == "" {
		log.Fatal("ELABS_API must be set")
	}
	log.Printf("Client Initialized\n")
	client := elevenlabs.NewClient(context.Background(), elevenAPI, 15*time.Second)
	return client

}

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
