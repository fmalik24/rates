package main

import (
	"strconv"
	"time"
)

func convertDayToNumber(day string) int {
	dayToNumber := map[string]int{
		"mon":   1,
		"tues":  2,
		"wed":   3,
		"thurs": 4,
		"fri":   5,
		"sat":   6,
		"sun":   0,
	}
	return dayToNumber[day]
}

func getNextDayNumber(number int) int {

	if number < 0 {
		return -1
	}
	if number+1 == 7 {
		return 0
	}
	return number + 1
}

func getPreviousDayNumber(number int) int {

	if number < 0 {
		return -1
	}

	if number == 0 {
		return 6
	}

	return number - 1
}

func convertNumberToDay(number int) string {
	numberToDay := map[int]string{
		1: "mon",
		2: "tues",
		3: "wed",
		4: "thurs",
		5: "fri",
		6: "sat",
		0: "sun",
	}
	return numberToDay[number]
}

func convertToTime(time time.Time) string {
	hour := strconv.Itoa(time.Hour())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	minutes := strconv.Itoa(time.Minute())
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	seconds := strconv.Itoa(time.Second())
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}
	return hour + minutes
}

func getDay(time time.Time) string {
	return convertNumberToDay(int(time.Weekday()))
}

func getDateTimeInUTC(day string, hours int, minutes int, location *time.Location) time.Time {

	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hours, minutes, 0, 0, location).UTC()
}

func getDateTime(day string, hours int, minutes int, location *time.Location) time.Time {

	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hours, minutes, 0, 0, location)
}
