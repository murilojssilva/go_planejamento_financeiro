package util

import "time"

const layout = "2006-01-02T15:05:05"

// StringToTime convert String to time
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)

	return convertedTime
}
