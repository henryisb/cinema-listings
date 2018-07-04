package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henryisb/cinema-listings/pkg/services"
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

	cinemas := services.GetCinemasByPostcode(postcode) //getDummyCinema(postcode)

	showingResponse := ShowingResponse{}

	// set response for showings
	for _, v := range cinemas.Cinemas {
		if v.Distance < 10.0 {
			showings := services.GetShowingsByCinemaID(v.ID, v.Name) //getDummyShowings()
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
// type Cinema struct {
// 	Postcode string `json:"postcode"`
// 	Cinemas  []struct {
// 		Name     string  `json:"name"`
// 		ID       string  `json:"id"`
// 		Distance float64 `json:"distance"`
// 	} `json:"cinemas"`
// }

// // Showings structure for each cinema
// type Showings struct {
// 	Status   string    `json:"status"`
// 	Listings []Listing `json:"listings"`
// }

// // Listing for each cinema
// type Listing struct {
// 	Title string
// 	Times []string
// }

// ShowingList for the cinema and listing
type ShowingList struct {
	Name    string
	Listing []services.Listing
}

// ShowingResponse to send back with list of listing
type ShowingResponse struct {
	ShowingList []ShowingList
}

// AddListing for appending the to the list of showings for response
func (showingResponse *ShowingResponse) AddListing(listing []services.Listing, name string) []ShowingList {
	showingEntity := ShowingList{name, listing}
	showingResponse.ShowingList = append(showingResponse.ShowingList, showingEntity)
	return showingResponse.ShowingList
}

// func getCinemasByPostcode(postcode string) Cinema {

// 	url := "http://api.cinelist.co.uk/search/cinemas/postcode/" + postcode

// 	reqClient := http.Client{
// 		Timeout: time.Second * 2,
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url, nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//req.Header.Set("User-Agent", "Cinema-listings")

// 	res, getErr := reqClient.Do(req)

// 	if getErr != nil {
// 		log.Fatal(getErr)
// 	}

// 	body, readErr := ioutil.ReadAll(res.Body)

// 	if readErr != nil {
// 		log.Fatal(readErr)
// 	}

// 	cinemas := Cinema{}
// 	json.Unmarshal(body, &cinemas)

// 	return cinemas

// }

// //GetShowingsByCinemaID brings back the showings for given cinema id
// func GetShowingsByCinemaID(id string, name string) Showings {

// 	url := "http://api.cinelist.co.uk/get/times/cinema/" + id

// 	reqClient := http.Client{
// 		Timeout: time.Second * 2,
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url, nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//req.Header.Set("User-Agent", "Cinema-listings")

// 	res, getErr := reqClient.Do(req)

// 	if getErr != nil {
// 		log.Fatal(getErr)
// 	}

// 	body, readErr := ioutil.ReadAll(res.Body)

// 	if readErr != nil {
// 		log.Fatal(readErr)
// 	}

// 	showings := Showings{}
// 	json.Unmarshal(body, &showings)

// 	return showings

// }

// // GetDummyShowings for testing
// func GetDummyShowings() Showings {

// 	dummy := []byte(`{
// 		"status": "ok",
// 		"listings": [
// 			{
// 				"title": "Jurassic World: Fallen Kingdom",
// 				"times": [
// 					"17:50"
// 				]
// 			},
// 			{
// 				"title": "In The Fade",
// 				"times": [
// 					"20:50"
// 				]
// 			},
// 			{
// 				"title": "Ocean's 8",
// 				"times": [
// 					"12:55",
// 					"15:25",
// 					"18:00"
// 				]
// 			},
// 			{
// 				"title": "Adrift",
// 				"times": [
// 					"11:00",
// 					"13:25",
// 					"15:45",
// 					"18:10"
// 				]
// 			},
// 			{
// 				"title": "Hereditary",
// 				"times": [
// 					"20:40"
// 				]
// 			},
// 			{
// 				"title": "The Happy Prince",
// 				"times": [
// 					"12:45"
// 				]
// 			},
// 			{
// 				"title": "Leave No Trace",
// 				"times": [
// 					"15:15",
// 					"20:30"
// 				]
// 			}
// 		]
// 	}`)

// 	showings := Showings{}
// 	json.Unmarshal(dummy, &showings)
// 	return showings

// }

// func getDummyCinema(postcode string) Cinema {
// 	dummy := []byte(`{
// 		"postcode": "yo264wy",
// 		"cinemas": [
// 			{
// 				"name": "City Screen Picturehouse, York",
// 				"id": "4370",
// 				"distance": 0.66
// 			}
// 		]
// 	}`)

// 	cinema := Cinema{}
// 	json.Unmarshal(dummy, &cinema)
// 	return cinema
// }

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
