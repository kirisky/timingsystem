package utils

import (
	"time"
	"timingsystem/timingserver/cerror"
)

// IsEmptyString check if the string is empty or not
func IsEmptyString(str string) bool {
	if len(str) > 0 {
		return false
	}

	return true
}

// GetUnixTime convert string date to timestamp
func GetUnixTime(timeString string) int64 {
	format := "2006-01-02 15:04:05"

	r, err := time.Parse(format, timeString)
	cerror.CheckErr(err)
	return r.Unix()
}

// MaxTime compare two times
// if atime is bigger, return 1
// if btime is bigger, return -1
// if they are the same, return 0
func MaxTime(atime int64, btime int64) int {

	switch {
	case atime > btime:
		return 1
	case atime < btime:
		return -1
	default:
		return 0
	}
}