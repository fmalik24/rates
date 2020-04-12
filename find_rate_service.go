package main

import (
	"fmt"
	"strconv"
	"time"
)

func findPrice(startDate string, endDate string) string {

	userStartDate, _ := time.Parse(time.RFC3339, startDate)
	userEndDate, _ := time.Parse(time.RFC3339, endDate)

	fmt.Printf("The start date is %s \n", userStartDate.String())
	fmt.Printf("The end date is %s \n", userEndDate.String())

	price := "unavailable"

	// Check if user is slot requests spans over a day
	if userEndDate.Day()-userStartDate.Day() >= 1 {
		return price
	}

	userStartDayUTC := convertNumberToDay(int(userStartDate.UTC().Weekday()))
	userEndDayUTC := convertNumberToDay(int(userEndDate.UTC().Weekday()))

	userStartTimeUTC := convertToTime(userStartDate.UTC())
	userEndTimeUTC := convertToTime(userEndDate.UTC())

	//Edge Case: Get the previous day as well to handle case where user start time utc is next day to slot start time
	//	         utc and user end time utc is less than slot end time utc.
	previousDay := convertNumberToDay(getPreviousDayNumber(convertDayToNumber(userStartDayUTC)))

	daySchedule := globalScheduleInUTC[userStartDayUTC]
	previousDaySchedule := globalScheduleInUTC[previousDay]

	twoDaySchedule := append(daySchedule, previousDaySchedule...)

	rateQueryFields := RateQueryFields{
		slots:            twoDaySchedule,
		userStartDayUTC:  userStartDayUTC,
		userEndDayUTC:    userEndDayUTC,
		userStartTimeUTC: userStartTimeUTC,
		userEndTimeUTC:   userEndTimeUTC,
	}

	price = queryPrice(rateQueryFields)

	return price
}

func queryPrice(rateQueryFields RateQueryFields) string {

	for i := 0; i < len(rateQueryFields.slots); i++ {

		userStartDayNumber := convertDayToNumber(rateQueryFields.userStartDayUTC)
		userEndDayNumber := convertDayToNumber(rateQueryFields.userEndDayUTC)

		slodStartDayNumber := convertDayToNumber(rateQueryFields.slots[i].startDay)
		slotEndDayNumber := convertDayToNumber(rateQueryFields.slots[i].endDay)

		userStartTime := rateQueryFields.userStartTimeUTC
		userEndTime := rateQueryFields.userEndTimeUTC

		slotStartTime := rateQueryFields.slots[i].startTime
		slotEndTime := rateQueryFields.slots[i].endTime

		price := rateQueryFields.slots[i].price

		// NOTE: User timings and slot timings are in UTC. Due to this a slot can span over two days but would be less than 24hr per contract

		// Case A
		// Slot spans over a day
		// User requested timings are with in the first day
		if getPreviousDayNumber(slotEndDayNumber) == userEndDayNumber && slodStartDayNumber == userEndDayNumber {
			if userStartTime >= slotStartTime && userEndTime > slotEndTime {
				return strconv.Itoa(price)
			}
		}

		// Case B
		// Slot spans over a day
		// User requested timing are in the second day
		if getPreviousDayNumber(userStartDayNumber) == slodStartDayNumber && slotEndDayNumber == userEndDayNumber {
			if userStartTime < slotStartTime && userEndTime <= slotEndTime {
				return strconv.Itoa(price)
			}
		}

		// Case C
		// Slots spans over a day
		// User requested start time is in first day and end time in the second day
		// OR
		// Slot is with in the same day
		// User requested start time and end time are in the same day
		if userStartDayNumber == slodStartDayNumber && userEndDayNumber == slotEndDayNumber {
			if userStartTime >= slotStartTime && userEndTime <= slotEndTime {
				return strconv.Itoa(price)
			}
		}
	}
	return "unavailable"

}
