package util

import (
	"time"
)

func GenerateSequenceNo() string {
	return time.Now().Format("20060102150405") + RandNumber(6)
}
