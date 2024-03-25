package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		timeUnix = time.Now().Unix()
	)
	fmt.Println("Current UNIX time is ", timeUnix)
	var timeSecond = timeUnix % 60
	var timeMinute = timeUnix / 60 % 60
	var timeHour = timeUnix / 3600 % 24
	fmt.Println("Current Time is ", timeHour, ":", timeMinute, ":", timeSecond, " GMT")
	var year int64 = 1970
	var month int64 = 1
	var dayinmonth int64 = 31
	var totalDAY int64
	for totalDAY = timeUnix / 86400; totalDAY >= dayinmonth; {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			totalDAY -= 31
			month++
			break
		case 4, 6, 9, 11:
			totalDAY -= 30
			month++
			break
		case 2:
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				totalDAY -= 29
			} else {
				totalDAY -= 28
			}
			month++
			break
		default:
			month = 1
			year++
		}
	}
	totalDAY += 1
	fmt.Println("Current date & time is", year, "-", month, "-", totalDAY, " ", timeHour, ":", timeMinute, ":", timeSecond, " GMT")
}
