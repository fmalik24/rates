package main

import (
	"fmt"
	"net/http"
)

var globalScheduleInUTC = map[string][]DailyScheduleUTC{}

func main() {
	http.HandleFunc("/getRates", getRates)
	http.ListenAndServe(":8080", nil)
}

func getRates(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		getRateForRequest(responseWriter, request)
	case "POST":
		saveRates(responseWriter, request)
	default:
		fmt.Fprint(responseWriter, "Sorry only get and post available")

	}

}

func getRateForRequest(responseWriter http.ResponseWriter, request *http.Request) {

	if len(globalScheduleInUTC) == 0 {
		repoInstance := repo{}
		globalScheduleInUTC = getScheduleUTC(repoInstance)
	}
	startTime := request.URL.Query()["startDate"][0]
	endTime := request.URL.Query()["endDate"][0]
	rate := findRate(startTime, endTime)
	fmt.Fprint(responseWriter, rate)
}
