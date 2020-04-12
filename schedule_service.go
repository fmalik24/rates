package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type DailyScheduleUTC struct {
	startTime string
	endTime   string
	endDay    string
	price     int
}

func getScheduleUTC(regPreCheck someHack) map[string][]DailyScheduleUTC {

	// first lets get the JSON
	rateBytes := regPreCheck.getDataFromFileSystem()
	rates := createRates(rateBytes)

	scheduleInUTC := make(map[string][]DailyScheduleUTC)
	createScheduleInUTC(rates, scheduleInUTC)

	return scheduleInUTC
}

func createScheduleInUTC(rates Rates, scheduleInUTC map[string][]DailyScheduleUTC) {
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

func createDaySchedule(rateEntry RateEntry, localDailyRatesUTC map[string][]DailyScheduleUTC) {
	for i := 0; i < len(rateEntry.days); i++ {
		populateDailySchedule(rateEntry, rateEntry.days[i], localDailyRatesUTC)
	}
}

func populateDailySchedule(rateInstance RateEntry, day string, localDailyRatesUTC map[string][]DailyScheduleUTC) {
	startTimeHours, _ := strconv.Atoi(rateInstance.startTime[0:2])
	endtimeHours, _ := strconv.Atoi(rateInstance.endTime[0:2])

	startMinutes, _ := strconv.Atoi(rateInstance.startTime[2:4])
	endMinutes, _ := strconv.Atoi(rateInstance.endTime[2:4])

	newStartTimeUTC := getDateTimeInUTC(day, startTimeHours, startMinutes, rateInstance.location)
	newEndTimeUTC := getDateTimeInUTC(day, endtimeHours, endMinutes, rateInstance.location)

	newStartTime := getDateTime(day, startTimeHours, startMinutes, rateInstance.location)
	newEndTime := getDateTime(day, endtimeHours, endMinutes, rateInstance.location)
	startDayNumber := convertDayToNumber(day)
	endDayNumber := startDayNumber

	startDayNumber = getDayNumberAfterTimeZoneConversion(newStartTime, newStartTimeUTC, startDayNumber)
	endDayNumber = getDayNumberAfterTimeZoneConversion(newEndTime, newEndTimeUTC, endDayNumber)

	var dailyScheduleUTC DailyScheduleUTC

	dailyScheduleUTC.endDay = convertNumberToDay(endDayNumber)
	dailyScheduleUTC.startTime = convertToTime(newStartTimeUTC)
	dailyScheduleUTC.endTime = convertToTime(newEndTimeUTC)
	dailyScheduleUTC.price = rateInstance.price
	dayUTC := convertNumberToDay(startDayNumber)

	if _, ok := localDailyRatesUTC[dayUTC]; ok {
		localDailyRatesUTC[dayUTC] = append(localDailyRatesUTC[dayUTC], dailyScheduleUTC)
	} else {
		var dailySchedulesUTC []DailyScheduleUTC
		localDailyRatesUTC[dayUTC] = append(dailySchedulesUTC, dailyScheduleUTC)
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
