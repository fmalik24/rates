package main

import "testing"

func TestGetRatesSameTimeZone(testHelper *testing.T) {

	// Arrange:  verify that no error occurred in the process
	testData = `{
		"rates": [
			{
				"days": "mon,tues,thurs",
				"times": "0600-0900",
				"tz": "America/New_York",
				"Price": 2500
			},
			{
				"days": "fri,wed,sat,sun",
				"times": "0600-0900",
				"tz": "America/New_York",
				"Price": 1500
			}
		]
	}`

	mockRequestClient := &mockFileSystem{
		mockGetDataFromFileSystem: func() []byte {
			return []byte(testData)
		},
	}

	// Act
	globalScheduleInUTC = getScheduleUTC(mockRequestClient)

	price := findRate("2015-07-01T06:00:00-04:00", "2015-07-01T09:00:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	price1 := findRate("2015-07-01T05:59:00-04:00", "2015-07-01T09:00:00-04:00")
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	price2 := findRate("2015-07-01T06:00:00-04:00", "2015-07-01T09:01:00-04:00")
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]ScheduleUTC)
}

func TestGetRatesDifferentTimeZone(testHelper *testing.T) {

	// Arrange:  verify that no error occurred in the process
	testData = `{
		"rates": [
			{
				"days": "mon,tues,thurs",
				"times": "0600-0900",
				"tz": "America/Chicago",
				"Price": 2500
			},
			{
				"days": "fri,wed,sat,sun",
				"times": "0600-0900",
				"tz": "America/Chicago",
				"Price": 1500
			}
		]
	}`

	mockRequestClient := &mockFileSystem{
		mockGetDataFromFileSystem: func() []byte {
			return []byte(testData)
		},
	}

	// Act
	globalScheduleInUTC = getScheduleUTC(mockRequestClient)

	price := findRate("2015-07-01T07:00:00-04:00", "2015-07-01T10:00:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	price1 := findRate("2015-07-01T06:59:00-04:00", "2015-07-01T10:00:00-04:00")
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	price2 := findRate("2015-07-01T07:00:00-04:00", "2015-07-01T10:01:00-04:00")
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]ScheduleUTC)

}

func TestGetRatesDifferentTimeZoneDifferentDay(testHelper *testing.T) {

	// Arrange:  verify that no error occurred in the process
	testData = `{
		"rates": [
			{
				"days": "mon,tues,thurs",
				"times": "2300-2350",
				"tz": "America/New_York",
				"Price": 2500
			},
			{
				"days": "fri,wed,sat,sun",
				"times":"2300-2350",
				"tz": "America/New_York",
				"Price": 1500
			}
		]
	}`

	mockRequestClient := &mockFileSystem{
		mockGetDataFromFileSystem: func() []byte {
			return []byte(testData)
		},
	}

	// Act
	globalScheduleInUTC = getScheduleUTC(mockRequestClient)

	// Assert
	price := findRate("2015-07-01T23:00:00-04:00", "2015-07-01T23:30:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	// Act
	price1 := findRate("2015-07-01T22:59:00-04:00", "2015-07-01T23:30:00-04:00")

	// Assert
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// Act
	price2 := findRate("2015-07-01T23:00:00-04:00", "2015-07-01T23:51:00-04:00")
	// Assert
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]ScheduleUTC)

}
