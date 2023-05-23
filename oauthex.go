package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

func main() {
	// Load the configuration from a JSON file
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err.Error())
		return
	}

	// Step 1: Redirect the user to the OAuth provider's authorization page
	authURL := fmt.Sprintf("https://oauth-provider.com/authorize?client_id=%s&redirect_uri=%s&response_type=code", config.ClientID, config.RedirectURL)
	fmt.Printf("Open the following URL in your browser and grant access:\n%s\n", authURL)

	// Step 2: Handle the callback from the OAuth provider
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the authorization code from the query parameters
		code := r.URL.Query().Get("code")

		// Step 3: Exchange the authorization code for an access token
		tokenURL := "https://oauth-provider.com/token"
		data := url.Values{
			"grant_type":    {"authorization_code"},
			"client_id":     {config.ClientID},
			"client_secret": {config.ClientSecret},
			"code":          {code},
			"redirect_uri":  {config.RedirectURL},
		}

		resp, err := http.PostForm(tokenURL, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Step 4: Parse the response and extract the access token
		var tokenResponse struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
		}
		err = json.Unmarshal(body, &tokenResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Step 5: Use the access token to make API requests
		accessToken := tokenResponse.AccessToken
		apiURL := "https://api.example.com/endpoint"

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response from the API
		apiResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Use the API response as needed
		fmt.Printf("API Response: %s\n", apiResponse)

		// Return a response to the client
		w.Write([]byte("Authentication and API call successful!"))
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// Function to load the configuration from a JSON file
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
