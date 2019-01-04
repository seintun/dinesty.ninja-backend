package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Config struct for Yelp Development API URL and key
type Config struct {
	YelpAPI struct {
		YelpAPIUrl string `json:"YelpAPIUrl"`
		YelpAPIKey string `json:"YelpAPIKey"`
	} `json:"YelpAPI"`
}

// LoadConfiguration reads, decodes JSON in config file and create in Config struct
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

// getBizYelp will extract incoming POST request body for businessID and
// pass it as param to api.yelp with GET request & Auth Bearer API key
func getBizYelp(w http.ResponseWriter, r *http.Request) {
	// read r.Body, declare as interface struct with the value of unmarshal []byte, and convert into string with Sprintf for APIUrl
	bizBody, _ := ioutil.ReadAll(r.Body)
	var bizIDkey map[string]interface{}
	json.Unmarshal([]byte(bizBody), &bizIDkey)
	bizID := fmt.Sprintf("%v", bizIDkey["businessID"])

	// invoke LoadConfiguration with json file path to get Yelp API info
	config, _ := LoadConfiguration("config/config.json")
	yelpURL := "https://api.yelp.com/v3/businesses/" + bizID
	bearer := "Bearer " + config.YelpAPI.YelpAPIKey
	// Send request to api.yelp with businessID and Auth bearer APIKey
	request, _ := http.NewRequest("GET", yelpURL, nil)
	request.Header.Add("Authorization", bearer)

	// Error-handle if missing response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// read response.Body from api.yelp and convert from []byte to string to JSON as response with Header of application/json
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	stringJSON := string(data)
	rawJSON := json.RawMessage(stringJSON)
	w.Header().Set("content-type", "application/json")
	w.Write(rawJSON)
}
