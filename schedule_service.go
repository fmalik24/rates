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

	newStartTime := getDateTimeInUTC(day, startTimeHours, startMinutes, rateInstance.location)
	newEndTime := getDateTimeInUTC(day, endtimeHours, endMinutes, rateInstance.location)

	var dailyScheduleUTC DailyScheduleUTC

	dailyScheduleUTC.endDay = getDay(newEndTime)
	dailyScheduleUTC.startTime = convertToTime(newStartTime)
	dailyScheduleUTC.endTime = convertToTime(newEndTime)
	dailyScheduleUTC.price = rateInstance.price
	dayUTC := getDay(newStartTime)

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
