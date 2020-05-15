package utils

import (
	"time"
)

func GenerateTaskID(format ...string) string {
	if len(format) == 0 {
		return time.Now().Format("20060102150405")
	} else {
		return time.Now().Format(format[0])
	}

}
