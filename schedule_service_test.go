package main

type mockFileSystem struct {
	mockGetDataFromFileSystem func() []byte
}

var testData = ""

func (m *mockFileSystem) getDataFromFileSystem() []byte {
	if m.mockGetDataFromFileSystem != nil {
		return []byte(testData)
	}
	return []byte("123")
}

// /* Test postiive */
// func TestGetScheduleUTC(testHelper *testing.T) {

// 	// Arrange:  verify that no error occurred in the process
// 	testData = `{
// 		"rates": [
// 			{
// 				"days": "mon",
// 				"times": "0900-2100",
// 				"tz": "America/Chicago",
// 				"Price": 1
// 			}
// 		]
// 	}`

// 	mockRequestClient := &mockRequestContext{
// 		mockGetDataFromFileSystem: func() []byte {
// 			return []byte(testData)
// 		},
// 	}

// 	// Act
// 	data := getScheduleUTC(mockRequestClient)

// 	// Assert

// 	fmt.Printf("What is data %v", data["mon"])
// 	if data["mon"][0].endDay != "tues" {
// 		testHelper.Errorf("Not expecting this: %s", data["mon"][0].endDay)
// 	}
// 	if data["mon"][0].endTime != "0300" {
// 		testHelper.Errorf("Not expecting this: %s", data["mon"][0].endTime)
// 	}
// 	if data["mon"][0].startTime != "1500" {
// 		testHelper.Errorf("Not expecting this: %s", data["mon"][0].startTime)
// 	}
// 	if data["mon"][0].price != 1500 {
// 		testHelper.Errorf("Not expecting this: %d", data["mon"][0].price)
// 	}

// }

// func TestGetScheduleUTC1(testHelper *testing.T) {

// 	// Arrange:  verify that no error occurred in the process
// 	testData = `{
// 		"rates": [
// 			{
// 				"days": "mon,tues,thurs,fri,wed,sat,sun",
// 				"times": "0600-0900",
// 				"tz": "America/New_York",
// 				"Price": 1500
// 			}
// 		]
// 	}`

// 	mockRequestClient := &mockFileSystem{
// 		mockGetDataFromFileSystem: func() []byte {
// 			return []byte(testData)
// 		},
// 	}

// 	// Act
// 	globalScheduleInUTC = getScheduleUTC(mockRequestClient)

// 	price := findRate("2015-07-01T14:20:00-03:00", "2015-07-01T15:21:00-03:00")

// 	if price == "2000" {

// 	}

// 	if len(globalScheduleInUTC) > 0 {

// 	}
// 	// Assert

// }

// func createScheduleInUTC(rates Rates, scheduleInUTC map[string][]ScheduleUTC)
// func getScheduleEntry(rate Rate) RateEntry
// func createDaySchedule(rateEntry RateEntry, scheduleInUTC map[string][]ScheduleUTC)
// func loadingDailySchedule(rateInstance RateEntry, day string, scheduleUTC map[string][]ScheduleUTC)
// func createRates(byteValue []byte) Rates
// func getDayNumberAfterTimeZoneConversion(timeInGivenTz time.Time, timeInUTC time.Time, dayNumber int)
