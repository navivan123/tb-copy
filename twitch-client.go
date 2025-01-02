package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expiry"`
}

type TokenStore struct {
	Token   string `json:"token"`
	Expires string `json:"expiry"`
}

const clientID = "2eeuai1d8j81wci0beavpwsaan8hcq"

var tokenChan = make(chan string)

func getAuthToken() string {

	expDate, token := retrieveToken()
	if (expDate == time.Time{} || token == "") || time.Now().After(expDate) {
		serv := startServer()
		defer serv.Close()
		token = callOauth(clientID)
		storeToken(token)

		return token
	}

	return token
}

func callOauth(clientID string) string {

	tokenURL := fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=%s&redirect_uri=http://localhost:3000&scope=channel%%3Aread%%3Aredemptions", clientID)

	err := browser.OpenURL(tokenURL)
	if err != nil {
		log.Fatal("Opening Twitch Auth in Browser failed")
	}

	token := waitForToken()
	if token == "" {
		log.Fatal("Error Obtaining Token!")
	}
	fmt.Printf("Access Token Received: %s\n", token)

	return token
}

func storeToken(token string) {
	expDate := time.Now().Add(24 * 60 * time.Hour).Format(time.RFC3339)
	tokenToStore := &TokenStore{Token: token, Expires: expDate}

	jsonData, err := json.Marshal(tokenToStore)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile("token.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func retrieveToken() (time.Time, string) {
	jsonFile, err := os.Open("token.json")
	if err != nil {
		log.Fatal("No json token file")
	}
	defer jsonFile.Close()

	jsonBytes, _ := io.ReadAll(jsonFile)
	var jsonToken Token

	json.Unmarshal(jsonBytes, &jsonToken)
	return jsonToken.Expires, jsonToken.Token
}

func startServer() *http.Server {
	const srcPort = "3000"

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", listen)
	serveMux.HandleFunc("/token", storeTokenHandler)

	httpServ := http.Server{
		Handler: serveMux,
		Addr:    ":" + srcPort,
	}

	log.Printf("Listening on port %s\n", srcPort)
	go httpServ.ListenAndServe()
	return &httpServ
}

func endServer() {}

func listen(w http.ResponseWriter, r *http.Request) {

	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>OAuth Redirect</title>
			<script>
				const token = new URLSearchParams(window.location.hash.substring(1)).get('access_token');
				if (token) {
					fetch('/token', {
						method: 'POST',
						headers: { 'Content-Type': 'application/json' },
						body: JSON.stringify({ token })
					}).then(() => {
						document.body.innerText = 'Token received! You can close this window.';
					}).catch(err => {
						console.error('Error:', err);
						document.body.innerText = 'Error sending token to the application.';
					});
				} else {
					document.body.innerText = 'No access token found. Something went wrong.';
				}
			</script>
		</head>
		<body>
			Redirecting...
		</body>
		</html>
	`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func storeTokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AAAAAAAAAAAAAAAAAAA")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data struct {
		Token string `json:"token"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if data.Token == "" {
		http.Error(w, "Empty token", http.StatusBadRequest)
		return
	}

	// Store the token in a global variable
	tokenChan <- data.Token

	fmt.Println("Access token received:", data.Token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token stored successfully"))
}

func waitForToken() string {
	select {
	case token := <-tokenChan:
		return token
	case <-time.After(5 * time.Minute):
		log.Fatal("Timeout waiting for token")
		return ""
	}
}
