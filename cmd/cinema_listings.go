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

// 	group by film?
//	showingresp -
//  	showinglist
// 			getName
//					(set object name)
//			listing
//				title
//					(set times =)
//

// ShowingByFilm gives the showings that can be presented by per film
type ShowingByFilm struct {
	films []Films
}

//Films that have timings for each cinema
type Films struct {
	name        string
	cinemaTimes []struct {
		name  string
		Times []string
	}
}

func (showingByFilm *ShowingByFilm) addShowing(showingResponse ShowingResponse) []Films {

	// for _, v := showingResponse.ShowingList {
	// 	name := v.name
	// 	for _, a := v.Listing {
	// 		a.Times
	// 	}
	// }

	return nil
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
