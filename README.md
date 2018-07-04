# Cinema Listings

## TODO
 - Search by film and location
		- brings back all the locations within 10 miles and times for single film
		- then add in day choosing
	- Choose distance to search within?
	- different days?
	- Non-code changes
		- UI? What to do the ui in?
		- Maybe make it app?
		- What to host it in? Serverless? Need to update the structure etc...
		- ...

give postcode (optional distance default 5?)
 - get all ids within n miles GET api.cinelist.co.uk/search/cinemas/postcode/:postcode
    - create object list by id and name of cinema

 for each id GET api.cinelist.co.uk/get/times/cinema/:venueID?day=<INT>
  - do get on venueId GET api.cinelist.co.uk/get/times/cinema/:venueID?day=<INT>
    - group together by title TODO
      - return
        - title : []
          - venue.name : []
            -  times
