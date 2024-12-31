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

func getVoices() []string {
	return []string{"Xb7hH8MSUJpSbSDYk0k2", "iP95p4xoKVk53GoZ742B", "pqHfZKP75CvOlQylNhV4"}
}

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

func callEleven(client *elevenlabs.Client, voice int, text string) {
	log.Printf("Making Request Struct\n")
	ttsReq := elevenlabs.TextToSpeechRequest{
		Text:    text,
		ModelID: modelID}

	log.Printf("Calling client\n")
	audio, err := client.TextToSpeech(getVoices()[voice], ttsReq)
	if err != nil {
		log.Printf("ERROR CALLING CLIENT\n")
		log.Fatal(err)
	}

	log.Printf("Writing File\n")
	if err = os.WriteFile("/tmp/elabs.mp3", audio, 0644); err != nil {
		log.Printf("ERROR WRITING TO FILE\n")
		log.Fatal(err)
	}

	log.Printf("Executing File\n")
	cmd := exec.Command("mpg123", "/tmp/elabs.mp3")
	if err = cmd.Start(); err != nil {
		log.Printf("ERROR EXECUTING FILE\n")
		log.Fatal(err)
	}
	if err = cmd.Wait(); err != nil {
		log.Printf("ERROR WAITING FOR COMMAND FILE")
		log.Fatal(err)
	}

}
