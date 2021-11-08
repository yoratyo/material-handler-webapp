package shared

import (
	"fmt"
	"time"
)

func GetCurrentTime() (string, string) {
	now := time.Now()
	y, m, d := now.Date()
	currDate := fmt.Sprintf("%d-%d-%d", y, int(m), d)
	currTime := fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second())

	return currDate, currTime
}
