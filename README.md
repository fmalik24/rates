# **Rates API**
------------

##  Description
   1.  API that allows a user to enter a date time range
    and get back the price at which they would be charged to park for that time span.
   2. API that allows to upload a new rate json 
   3. API that provides performance metrics
    Note: Find API details in swagger.yml
    
## User Guide
    
    1. Clone the repo
    2. cd in the project directory i.e /rates
    
    Docker 
    1. docker build -t rates .
    2. docker run -p 8080:8080 rates

    Local System
    1. go build ../rates
    2. go run ../rates

    Tests
    1. go test
    2. go test -coverprofile cp.out 


## Sample JSON for testing
```json
{
    "rates": [
        {
            "days": "mon,tues,thurs",
            "times": "0900-2100",
            "tz": "America/Chicago",
            "price": 1500
        },
        {
            "days": "fri,sat,sun",
            "times": "0900-2100",
            "tz": "America/Chicago",
            "price": 2000
        },
        {
            "days": "wed",
            "times": "0600-1800",
            "tz": "America/Chicago",
            "price": 1750
        },
        {
            "days": "mon,wed,sat",
            "times": "0100-0500",
            "tz": "America/Chicago",
            "price": 1000
        },
        {
            "days": "sun,tues",
            "times": "0100-0700",
            "tz": "America/Chicago",
            "price": 925
        }
    ]
}

```
The timezones specified in the JSON file adhere to the 2017c version of the tz database.

## Sample result
Datetime ranges should be specified in ISO-8601 format.


* `2015-07-01T07:00:00-05:00` to `2015-07-01T12:00:00-05:00`  yields `1750`
* http://localhost:8080/rates?startDate=2015-07-01T07:00:00-05:00&endDate=2015-07-01T12:00:00-05:00

* `2015-07-04T15:00:00+00:00` to `2015-07-04T20:00:00+00:00`  yields `2000`
* http://localhost:8080/rates?startDate=2015-07-04T15:00:00+00:00&endDate=2015-07-04T20:00:00+00:00

* `2015-07-04T07:00:00+05:00` to `2015-07-04T20:00:00+05:00`  yields `unavailable`
* http://localhost:8080/rates?startDate=2015-07-04T07:00:00+05:00&endDate=2015-07-04T20:00:00+05:00