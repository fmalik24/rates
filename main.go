package main

import (
	"fmt"
	"net/http"
	"strings"
)

var globalScheduleInUTC = map[string][]SlotUTC{}

func main() {
	http.HandleFunc("/rates", ratesAPI)
	http.ListenAndServe(":8080", nil)
}

func ratesAPI(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		getRateForRequest(responseWriter, request)
	case "POST":
		saveRates(responseWriter, request)
	default:
		fmt.Fprint(responseWriter, "Sorry only get and post available")

	}

}

/**
getRateForRequest does two things
1. Converts the rates json to a map of days having its schedule as itarray in UTC.
   Note:

*/

func getRateForRequest(responseWriter http.ResponseWriter, request *http.Request) {

	// Check if gloablScheduleInUTC has been populated with the schedule
	if len(globalScheduleInUTC) == 0 {
		repoInstance := repo{}
		globalScheduleInUTC = getScheduleUTC(repoInstance)
	}
	startTime := request.URL.Query()["startDate"][0]
	endTime := request.URL.Query()["endDate"][0]

	// Edge Case: + in url is reservered for " ". API is king!
	startTime = strings.Replace(startTime, " ", "+", -1)
	endTime = strings.Replace(endTime, " ", "+", -1)

	rate := findPrice(startTime, endTime)
	fmt.Fprint(responseWriter, rate)
}
