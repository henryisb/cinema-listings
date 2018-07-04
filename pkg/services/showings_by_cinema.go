package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Showings structure for each cinema
type Showings struct {
	Status   string    `json:"status"`
	Listings []Listing `json:"listings"`
}

// Listing for each cinema
type Listing struct {
	Title string
	Times []string
}

//GetShowingsByCinemaID brings back the showings for given cinema id
func GetShowingsByCinemaID(id string, name string) Showings {

	url := "http://api.cinelist.co.uk/get/times/cinema/" + id

	reqClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "Cinema-listings")

	res, getErr := reqClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	showings := Showings{}
	json.Unmarshal(body, &showings)

	return showings

}

// GetDummyShowings for testing
func GetDummyShowings() Showings {

	dummy := []byte(`{
		"status": "ok",
		"listings": [
			{
				"title": "Jurassic World: Fallen Kingdom",
				"times": [
					"17:50"
				]
			},
			{
				"title": "In The Fade",
				"times": [
					"20:50"
				]
			},
			{
				"title": "Ocean's 8",
				"times": [
					"12:55",
					"15:25",
					"18:00"
				]
			},
			{
				"title": "Adrift",
				"times": [
					"11:00",
					"13:25",
					"15:45",
					"18:10"
				]
			},
			{
				"title": "Hereditary",
				"times": [
					"20:40"
				]
			},
			{
				"title": "The Happy Prince",
				"times": [
					"12:45"
				]
			},
			{
				"title": "Leave No Trace",
				"times": [
					"15:15",
					"20:30"
				]
			}
		]
	}`)

	showings := Showings{}
	json.Unmarshal(dummy, &showings)
	return showings

}
