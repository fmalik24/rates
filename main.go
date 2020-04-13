package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var globalScheduleInUTC = map[string][]SlotUTC{}
var executionTimes []time.Duration

func main() {
	http.HandleFunc("/rates", ratesAPI)
	http.HandleFunc("/metrics", getMetrics)
	http.ListenAndServe(":8080", nil)
}

func ratesAPI(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		executionStartTime := time.Now()
		getRateForRequest(responseWriter, request)
		executionTimes = append(executionTimes, time.Now().Sub(executionStartTime))

	case "POST":
		executionStartTime := time.Now()
		saveRates(responseWriter, request)
		executionTimes = append(executionTimes, time.Now().Sub(executionStartTime))
	default:
		fmt.Fprint(responseWriter, "Sorry only get and post available")

	}

}

func getMetrics(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		var sum int64
		len := len(executionTimes)
		if len > 0 {
			for i := 0; i < len; i++ {
				sum += executionTimes[i].Microseconds()
			}
			result := sum / int64(len)
			fmt.Fprintf(responseWriter, "Mean Time %d ms\n", result)
			min, max := findMinAndMax(executionTimes)
			fmt.Fprintf(responseWriter, "Min Time %d ms \nMax Time %d ms\nRequest count %d\n", min, max, len)
		}

	}
}

func findMinAndMax(a []time.Duration) (min int64, max int64) {
	min = a[0].Microseconds()
	max = a[0].Microseconds()
	for _, value := range a {
		if value.Microseconds() < min {
			min = value.Microseconds()
		}
		if value.Microseconds() > max {
			max = value.Microseconds()
		}
	}
	return min, max
}

/**
getRateForRequest does two things
1. Converts the rates json to a map of days having its schedule in UTC in globalScheduleInUTC
2. Find the price in the globalScheduleInUTC for the given user time range
*/

func getRateForRequest(responseWriter http.ResponseWriter, request *http.Request) {

	// Check if gloablScheduleInUTC has been populated with the schedule
	if len(globalScheduleInUTC) == 0 {
		repoInstance := repo{}
		globalScheduleInUTC = getScheduleUTC(repoInstance)
	}
	startTime := request.URL.Query()["startDate"]
	endTime := request.URL.Query()["endDate"]

	if startTime == nil || endTime == nil {
		http.Error(responseWriter, "Bad Inputs", http.StatusBadRequest)
		return
	}

	startTime[0] = strings.Replace(startTime[0], " ", "+", -1)
	endTime[0] = strings.Replace(endTime[0], " ", "+", -1)

	rate := findPrice(startTime[0], endTime[0])
	fmt.Fprint(responseWriter, rate)
}
