package main

import (
	"fmt"
	"strconv"
	"time"
)

func findRate(startDate string, endDate string) string {

	userStartDate, _ := time.Parse(time.RFC3339, startDate)
	userEndDate, _ := time.Parse(time.RFC3339, endDate)

	fmt.Printf("The start date is %s", userStartDate.String())
	fmt.Printf("The start date is %s", userEndDate.String())

	result := "unavailable"

	// Check if user is slot requests spans over a day
	if userEndDate.Day()-userStartDate.Day() >= 1 {
		return result
	}

	userStartDayUTC := convertNumberToDay(int(userStartDate.UTC().Weekday()))
	userEndDayUTC := convertNumberToDay(int(userEndDate.UTC().Weekday()))

	userStartTimeUTC := convertToTime(userStartDate.UTC())
	userEndTimeUTC := convertToTime(userEndDate.UTC())

	daySchedule := globalScheduleInUTC[userStartDayUTC]
	result = findRatesFromScheudle(daySchedule, userEndDayUTC, userStartTimeUTC, userEndTimeUTC)

	fmt.Printf("The rate given is %s \n", result)
	return result
}

func findRatesFromScheudle(daySchedule []DailyScheduleUTC, userEndDayUTC string, userStartTimeUTC string, userEndTimeUTC string) string {
	for i := 0; i < len(daySchedule); i++ {
		if convertDayToNumber(daySchedule[i].endDay) == getNextDayNumber(convertDayToNumber(userEndDayUTC)) {
			if daySchedule[i].startTime <= userStartTimeUTC && daySchedule[i].endTime < userEndTimeUTC {
				return strconv.Itoa(daySchedule[i].price)
			}
		} else if daySchedule[i].startTime <= userStartTimeUTC && daySchedule[i].endTime >= userEndTimeUTC {
			return strconv.Itoa(daySchedule[i].price)
		}
	}
	return "unavailable"
}
