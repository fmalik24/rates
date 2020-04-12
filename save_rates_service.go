package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func saveRates(responseWriter http.ResponseWriter, request *http.Request) {
	var rates Rates
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(request.Body).Decode(&rates)
	if err != nil {
		/*
			TODO: // Important Remianing Work GET
		*/
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	ratesJSON, _ := json.MarshalIndent(rates, "", "    ")
	fmt.Printf("The Json is %s", string(ratesJSON))
	// Do something with the Person struct...
	fmt.Fprintf(responseWriter, "Rates: %+v", rates)

	ratesFile, err := os.Create("user.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = ratesFile.Write(ratesJSON)
	if err != nil {
		fmt.Println(err)
		ratesFile.Close()
		return
	}

	err = ratesFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSystem := repo{}
	globalScheduleInUTC = getScheduleUTC(fileSystem)

}
