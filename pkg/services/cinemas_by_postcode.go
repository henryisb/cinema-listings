package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Cinema type for list of responses to get cinemas by postcode
type Cinema struct {
	Postcode string `json:"postcode"`
	Cinemas  []struct {
		Name     string  `json:"name"`
		ID       string  `json:"id"`
		Distance float64 `json:"distance"`
	} `json:"cinemas"`
}

func GetCinemasByPostcode(postcode string) Cinema {

	url := "http://api.cinelist.co.uk/search/cinemas/postcode/" + postcode

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

	cinemas := Cinema{}
	json.Unmarshal(body, &cinemas)

	return cinemas

}

// GetDummyCinema For testing returned dummy data
func GetDummyCinema(postcode string) Cinema {
	dummy := []byte(`{
		"postcode": "yo264wy",
		"cinemas": [
			{
				"name": "City Screen Picturehouse, York",
				"id": "4370",
				"distance": 0.66
			}
		]
	}`)

	cinema := Cinema{}
	json.Unmarshal(dummy, &cinema)
	return cinema
}
