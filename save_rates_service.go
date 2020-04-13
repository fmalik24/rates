package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func saveRates(responseWriter http.ResponseWriter, request *http.Request) {
	var rates Rates
	err := json.NewDecoder(request.Body).Decode(&rates)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	ratesJSON, _ := json.MarshalIndent(rates, "", "    ")

	fileSystem := repo{}

	fileSystem.saveDataToFileSystem(ratesJSON)
	globalScheduleInUTC = getScheduleUTC(fileSystem)

	fmt.Fprintf(responseWriter, "Rates: %+v", rates)
}
