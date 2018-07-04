package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := newRouter()

	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")

}

func handlerCinemas(w http.ResponseWriter, r *http.Request) {
	// get postcode from path
	pathVar := mux.Vars(r)
	postcode := pathVar["postcode"]

	cinemas := getCinemasByPostcode(postcode)

	showingResponse := ShowingResponse{}

	// set response for showings
	for _, v := range cinemas.Cinemas {
		if v.Distance < 10.0 {
			showings := getShowingsByCinemaID(v.ID, v.Name)
			showingResponse.AddListing(showings.Listings, v.Name)
		}
	}

	// map to json
	jsonResp := json.NewEncoder(w).Encode(showingResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, jsonResp)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/cinemas/{postcode}", handlerCinemas).Methods("GET")

	return r
}

// Cinema type for list of responses to get cinemas by postcode
type Cinema struct {
	Postcode string `json:"postcode"`
	Cinemas  []struct {
		Name     string  `json:"name"`
		ID       string  `json:"id"`
		Distance float64 `json:"distance"`
	} `json:"cinemas"`
}

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

// ShowingList for the cinema and listing
type ShowingList struct {
	Name    string
	Listing []Listing
}

// ShowingResponse to send back with list of listing
type ShowingResponse struct {
	ShowingList []ShowingList
}

// AddListing for appending the to the list of showings for response
func (showingResponse *ShowingResponse) AddListing(listing []Listing, name string) []ShowingList {
	showingEntity := ShowingList{name, listing}
	showingResponse.ShowingList = append(showingResponse.ShowingList, showingEntity)
	return showingResponse.ShowingList
}

func getCinemasByPostcode(postcode string) Cinema {

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

func getShowingsByCinemaID(id string, name string) Showings {

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

//TODO
//  - Search by film and location
//		- brings back all the locations within 10 miles and times for single film
//		- then add in day choosing
//	- Choose distance to search within?
//	- different days?
//	- Non-code changes
//		- UI? What to do the ui in?
//		- Maybe make it app?
//		- What to host it in? Serverless? Need to update the structure etc...
//		-

// give postcode (optional distance default 5?)
//  - get all ids within n miles?  GETapi.cinelist.co.uk/search/cinemas/postcode/:postcode
//     - create object list by id and name of cinema

//  -for each id GETapi.cinelist.co.uk/get/times/cinema/:venueID?day=<INT>
//   - do get on venueId GETapi.cinelist.co.uk/get/times/cinema/:venueID?day=<INT>
//     - group together by title
//     - return
//      - title : []
//       - venue.name : []
//        - times

// sample
// {
//   "postcode": "yo264wy",
//   "cinemas": [
//     {
//       "name": "City Screen Picturehouse, York",
//       "id": "4370",
//       "distance": 0.66
//     },
//     {
//       "name": "VUE York, Rawcliffe",
//       "id": "9467",
//       "distance": 2.1
//     },
//     {
//       "name": "Wetherby Film Theatre, North Wetherby",
//       "id": "10660",
//       "distance": 11.9
//     },
//     {
//       "name": "Odeon Harrogate, Harrogate",
//       "id": "9396",
//       "distance": 17.78
//     },
//     {
//       "name": "Cineworld Castleford, Castleford",
//       "id": "7506",
//       "distance": 19.09
//     }
//   ]
// }
