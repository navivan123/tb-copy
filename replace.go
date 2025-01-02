package main

import (
	//"fmt"
	"strconv"
	"strings"
)

const peopleSize = 3

func getVoices() []string {
	return []string{"Xb7hH8MSUJpSbSDYk0k2", "iP95p4xoKVk53GoZ742B", "pqHfZKP75CvOlQylNhV4"}
}

func getPeople() []string {
	return []string{"alice:", "chris:", "bill:"}
}

func replace(text string) []string {

	s := strings.ToLower(text)
	arr := r_split(s, 0)
	arr = arr[1:]

	return arr
}

func r_split(text string, idx int) []string {
	arr := strings.Split(text, getPeople()[idx])
	var arr2 []string
	if idx == (peopleSize)-1 {
		for i, s := range arr {
			if i != 0 {
				s = strconv.Itoa(idx) + s
			}
			arr2 = append(arr2, s)
		}
		return arr2
	}
	for i, s := range arr {
		if i != 0 {
			s = strconv.Itoa(idx) + s
		}
		arr2 = append(arr2, r_split(s, idx+1)...)
	}
	return arr2
}

func initiateElevenLabs(inputText string) {
	client := initEleven()
	arr := replace(inputText)
	for _, voice := range arr {
		idx, _ := strconv.Atoi(voice[0:1])
		callEleven(client, getVoices()[idx], voice[1:])
	}
}
