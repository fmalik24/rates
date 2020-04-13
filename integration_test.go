package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Important: These are integration tests. They are dealing with an actual file system.
// Tests would manipulate the rate.json.

func TestGetRate(testHelper *testing.T) {

	handlerFunction := http.HandlerFunc(ratesAPI)
	url := "/rates?startDate=2015-07-04T15:00:00+00:00&endDate=2015-07-04T20:00:00+00:00"
	httpVerb := "GET"

	// Setup the request
	request, err := http.NewRequest(httpVerb, url, nil)
	if err != nil {
		testHelper.Fatal(err)
	}

	// Setup the response recorder
	response := httptest.NewRecorder()

	//Act:
	// Trigger HTTP request with the given data
	handlerFunction.ServeHTTP(response, request)

	// Assert:
	// The status code is as per expectation
	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPostRate(testHelper *testing.T) {

	sampleJSON := `{
		"rates": [
			{
				"days": "mon,tues,thurs",
				"times": "0900-2100",
				"tz": "America/Chicago",
				"price": 1500
			},
			{
				"days": "fri,sat,sun",
				"times": "0900-2100",
				"tz": "America/Chicago",
				"price": 2000
			},
			{
				"days": "wed",
				"times": "0600-1800",
				"tz": "America/Chicago",
				"price": 1750
			},
			{
				"days": "mon,wed,sat",
				"times": "0100-0500",
				"tz": "America/Chicago",
				"price": 1000
			},
			{
				"days": "sun,tues",
				"times": "0100-0700",
				"tz": "America/Chicago",
				"price": 925
			}
		]
	}`
	handlerFunction := http.HandlerFunc(ratesAPI)
	url := "/rates"
	httpVerb := "POST"

	// Setup the request
	request, err := http.NewRequest(httpVerb, url, bytes.NewReader([]byte(sampleJSON)))
	if err != nil {
		testHelper.Fatal(err)
	}

	// Setup the Content-Type to be of MultipartFomData
	request.Header.Set("Content-Type", "application/json")

	// Setup the response recorder
	response := httptest.NewRecorder()

	//Act:
	// Trigger HTTP request with the given data
	handlerFunction.ServeHTTP(response, request)

	// Assert:
	// The status code is as per expectation
	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
