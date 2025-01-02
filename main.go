package main

import (
	"fmt"
)

func main() {
	token := getAuthToken()
	pubsubRun(token)
	fmt.Println("Press any key or Ctrl+C to stop!")
	fmt.Scanln()
}
