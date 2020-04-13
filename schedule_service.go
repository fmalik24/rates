package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type SlotUTC struct {
	startTime string
	endTime   string
	endDay    string
	startDay  string
	price     int
}

func getScheduleUTC(fsys IFileSystem) map[string][]SlotUTC {

	// Lets get the JSON bytes from the file system
	rateBytes := fsys.getDataFromFileSystem()

	// Then get the rates Rates instance intialized with rateBytes
	rates := createRates(rateBytes)

	scheduleInUTC := make(map[string][]SlotUTC)

	// Create the Schedule in UTC
	createScheduleInUTC(rates, scheduleInUTC)

	return scheduleInUTC
}

func createScheduleInUTC(rates Rates, scheduleInUTC map[string][]SlotUTC) {

	// Loop thru each rate entry and update the scehdule by massaging the rate entry

	for i := 0; i < len(rates.Rates); i++ {
		scheduleInstance := getScheduleEntry(rates.Rates[i])
		createDaySchedule(scheduleInstance, scheduleInUTC)
	}
}

func getScheduleEntry(rate Rate) RateEntry {

	times := strings.Split(rate.Times, "-")
	location, _ := time.LoadLocation(rate.Tz)

	return RateEntry{
		days:      strings.Split(rate.Days, ","),
		startTime: times[0],
		endTime:   times[1],
		location:  location,
		price:     rate.Price,
	}

}

func createDaySchedule(rateEntry RateEntry, scheduleInUTC map[string][]SlotUTC) {

	// Loop thru each day and update the schedule
	for i := 0; i < len(rateEntry.days); i++ {
		loadingDailySchedule(rateEntry, rateEntry.days[i], scheduleInUTC)
	}
}

func loadingDailySchedule(rateInstance RateEntry, day string, scheduleUTC map[string][]SlotUTC) {

	// Getting the hours from rate entry
	startTimeHours, _ := strconv.Atoi(rateInstance.startTime[0:2])
	endtimeHours, _ := strconv.Atoi(rateInstance.endTime[0:2])

	// Getting the minutes from the rate entry
	startMinutes, _ := strconv.Atoi(rateInstance.startTime[2:4])
	endMinutes, _ := strconv.Atoi(rateInstance.endTime[2:4])

	// Converting rate entry time from its timezone to UTC
	newStartTimeUTC := getDateTimeInUTC(day, startTimeHours, startMinutes, rateInstance.location)
	newEndTimeUTC := getDateTimeInUTC(day, endtimeHours, endMinutes, rateInstance.location)

	// Edge Case: Check if after changing time zone the day in UTC changes or not.
	newStartTime := getDateTime(day, startTimeHours, startMinutes, rateInstance.location)
	newEndTime := getDateTime(day, endtimeHours, endMinutes, rateInstance.location)
	startDayNumber := convertDayToNumber(day)
	endDayNumber := startDayNumber

	// Get the day after converting to timezone
	startDayNumber = getDayNumberAfterTimeZoneConversion(newStartTime, newStartTimeUTC, startDayNumber)
	endDayNumber = getDayNumberAfterTimeZoneConversion(newEndTime, newEndTimeUTC, endDayNumber)
	dayUTC := convertNumberToDay(startDayNumber)

	var dailyScheduleUTC SlotUTC

	dailyScheduleUTC.endDay = convertNumberToDay(endDayNumber)
	dailyScheduleUTC.startTime = convertToTime(newStartTimeUTC)
	dailyScheduleUTC.endTime = convertToTime(newEndTimeUTC)
	dailyScheduleUTC.price = rateInstance.price
	dailyScheduleUTC.startDay = dayUTC

	//  scheduleUTC maps's key is the day of the UTC time
	if _, ok := scheduleUTC[dayUTC]; ok {
		scheduleUTC[dayUTC] = append(scheduleUTC[dayUTC], dailyScheduleUTC)
	} else {
		var schedulesUTC []SlotUTC
		scheduleUTC[dayUTC] = append(schedulesUTC, dailyScheduleUTC)
	}

}

func createRates(byteValue []byte) Rates {
	var rates Rates
	json.Unmarshal([]byte(byteValue), &rates)
	return rates
}

func getDayNumberAfterTimeZoneConversion(timeInGivenTz time.Time, timeInUTC time.Time, dayNumber int) int {
	if int(timeInGivenTz.Weekday()) < int(timeInUTC.Weekday()) {
		dayNumber = getNextDayNumber(dayNumber)
	}

	if int(timeInGivenTz.Weekday()) > int(timeInUTC.Weekday()) {
		dayNumber = getPreviousDayNumber(dayNumber)
	}

	return dayNumber

}
