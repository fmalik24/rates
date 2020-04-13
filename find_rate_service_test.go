package main

import (
	"testing"
)

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

	price := findPrice("2015-07-01T06:00:00-04:00", "2015-07-01T09:00:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	price1 := findPrice("2015-07-01T05:59:00-04:00", "2015-07-01T09:00:00-04:00")
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	price2 := findPrice("2015-07-01T06:00:00-04:00", "2015-07-01T09:01:00-04:00")
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]SlotUTC)
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

	price := findPrice("2015-07-01T07:00:00-04:00", "2015-07-01T10:00:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	price1 := findPrice("2015-07-01T06:59:00-04:00", "2015-07-01T10:00:00-04:00")
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	price2 := findPrice("2015-07-01T07:00:00-04:00", "2015-07-01T10:01:00-04:00")
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]SlotUTC)

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
	price := findPrice("2015-07-01T23:00:00-04:00", "2015-07-01T23:30:00-04:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	// Act
	price1 := findPrice("2015-07-01T22:59:00-04:00", "2015-07-01T23:30:00-04:00")

	// Assert
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// Act
	price2 := findPrice("2015-07-01T23:00:00-04:00", "2015-07-01T23:51:00-04:00")
	// Assert
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]SlotUTC)

}

func TestGetRatesDifferentTimeZoneDifferentDay1(testHelper *testing.T) {

	// Arrange:  verify that no error occurred in the process
	testData = `{
		"rates": [
			{
				"days": "mon,tues,thurs",
				"times": "2300-2350",
				"tz": "Asia/Karachi",
				"Price": 2500
			},
			{
				"days": "fri,wed,sat,sun",
				"times":"2300-2350",
				"tz": "Asia/Karachi",
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
	price := findPrice("2015-07-01T23:00:00+05:00", "2015-07-01T23:30:00+05:00")
	if price != "1500" {
		testHelper.Errorf("Got %s \nExpeted 1500", price)
	}

	// Act
	price1 := findPrice("2015-07-01T22:59:00+05:00", "2015-07-01T23:30:00+05:00")

	// Assert
	if price1 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price1)
	}

	// Act
	price2 := findPrice("2015-07-01T23:00:00+05:00", "2015-07-01T23:51:00+05:00")
	// Assert
	if price2 != "unavailable" {
		testHelper.Errorf("Got %s \nExpeted unavailable", price2)
	}

	// tear down
	globalScheduleInUTC = make(map[string][]SlotUTC)

}

func TestQueryPriceCaseC1(testHelper *testing.T) {

	// Slot is with in the same day
	// User requested start time and end time are in the same day

	slots := []SlotUTC{
		{
			startTime: "1000",
			endTime:   "1100",
			price:     1000,
			endDay:    "mon",
			startDay:  "mon",
		},
		{
			startTime: "0800",
			endTime:   "1000",
			price:     1,
			endDay:    "mon",
			startDay:  "mon",
		}}

	rateQueryFields := RateQueryFields{
		slots:            slots,
		userStartDayUTC:  "mon",
		userEndDayUTC:    "mon",
		userStartTimeUTC: "1000",
		userEndTimeUTC:   "1100",
	}

	price := queryPrice(rateQueryFields)
	if price != "1000" {
		testHelper.Errorf("Got %s \nExpeted 1000", price)
	}

}

func TestQueryPriceCaseC(testHelper *testing.T) {

	// Case C
	// Slots spans over a day
	// User requested start time is in first day and end time in the second day

	slots := []SlotUTC{
		{
			startTime: "1000",
			endTime:   "0900",
			price:     1000,
			endDay:    "tue",
			startDay:  "mon",
		},
		{
			startTime: "0800",
			endTime:   "1000",
			price:     1,
			endDay:    "mon",
			startDay:  "mon",
		}}

	rateQueryFields := RateQueryFields{
		slots:            slots,
		userStartDayUTC:  "mon",
		userEndDayUTC:    "tue",
		userStartTimeUTC: "1000",
		userEndTimeUTC:   "0900",
	}

	price := queryPrice(rateQueryFields)
	if price != "1000" {
		testHelper.Errorf("Got %s \nExpeted 1000", price)
	}

}

func TestQueryPriceCaseB(testHelper *testing.T) {

	// Case B
	// Slot spans over a day
	// User requested timing are in the second day

	slots := []SlotUTC{
		{
			startTime: "1100",
			endTime:   "0900",
			price:     1000,
			endDay:    "tues",
			startDay:  "mon",
		},
		{
			startTime: "1000",
			endTime:   "1100",
			price:     1,
			endDay:    "mon",
			startDay:  "mon",
		}}

	rateQueryFields := RateQueryFields{
		slots:            slots,
		userStartDayUTC:  "tues",
		userEndDayUTC:    "tues",
		userStartTimeUTC: "0800",
		userEndTimeUTC:   "0900",
	}

	price := queryPrice(rateQueryFields)
	if price != "1000" {
		testHelper.Errorf("Got %s \nExpeted 1000", price)
	}

}

func TestQueryPriceCaseA(testHelper *testing.T) {

	// Case A
	// Slot spans over a day
	// User requested timings are with in the first day

	slots := []SlotUTC{
		{
			startTime: "1100",
			endTime:   "0900",
			price:     1000,
			endDay:    "tues",
			startDay:  "mon",
		},
		{
			startTime: "1000",
			endTime:   "1100",
			price:     1,
			endDay:    "mon",
			startDay:  "mon",
		}}

	rateQueryFields := RateQueryFields{
		slots:            slots,
		userStartDayUTC:  "mon",
		userEndDayUTC:    "mon",
		userStartTimeUTC: "1100",
		userEndTimeUTC:   "1200",
	}

	price := queryPrice(rateQueryFields)
	if price != "1000" {
		testHelper.Errorf("Got %s \nExpeted 1000", price)
	}

}
