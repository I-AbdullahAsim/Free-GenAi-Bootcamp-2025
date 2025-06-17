package utils

import (
	"time"
)

// FormatDateTime formats a time.Time into a standard string format
func FormatDateTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// FormatDate formats a date into YYYY-MM-DD format
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTime formats a time into HH:MM:SS format
func FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}

// CalculateDateDifference calculates the difference between two dates in days
func CalculateDateDifference(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// CalculateTimeDifference calculates the difference between two times in hours
func CalculateTimeDifference(start, end time.Time) float64 {
	return end.Sub(start).Hours()
}

// IsSameDay checks if two times are on the same calendar day
func IsSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day()
}

// GetStartOfDay returns the start of the day (00:00:00) for a given time
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay returns the end of the day (23:59:59) for a given time
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
