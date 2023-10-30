package tools

import (
	"strconv"
	"time"
)

const dateTimeFormat = "2006-01-02 15:04:05"

func Now() string {
	return time.Now().Format(dateTimeFormat)
}

// The T is just a literal to separate the date from the time, and the Z means “zero hour offset” also known as “Zulu time” (UTC).
// 2023-09-27T11:07:36Z
func ZuluTime(dateString string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z", dateString)
	if err != nil {
		return ""
	}
	return t.Format(dateTimeFormat)
}

// RFC3339 2023-09-27T17:29:32+08:00
func RFC3339(dateString string) string {
	time, err := time.Parse("2006-01-02T15:04:05Z07:00", dateString)
	if err != nil {
		return ""
	}
	return time.Format(dateTimeFormat)
}

// 05月17日 10点15分
func NcnlTime(dateString string) string {
	year := time.Now().Year()
	dateString = strconv.Itoa(year) + "-" + dateString
	t, err := time.Parse("2006-01月02日 15点04分", dateString)
	if err != nil {
		return ""
	}
	return t.Format(dateTimeFormat)
}
