package main

import "time"

type RateEntry struct {
	days      []string
	startTime string
	endTime   string
	price     int
	location  *time.Location
}

// Rates struct which contains
// an array of users
type Rates struct {
	Rates []Rate `json:"rates"`
}

// Rate struct which contains a name
// a type and a list of social links
type Rate struct {
	Days  string `json:"days"`
	Times string `json:"times"`
	Tz    string `json:"tz"`
	Price int    `json:"Price"`
}

type RateQueryFields struct {
	slots            []SlotUTC
	userEndDayUTC    string
	userStartTimeUTC string
	userEndTimeUTC   string
	userStartDayUTC  string
}
