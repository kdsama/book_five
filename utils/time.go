package utils

import "time"

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
func OneHourInSeconds() int64 {
	return int64(3600)
}
