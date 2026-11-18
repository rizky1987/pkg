package utils

import (
	"time"
)

const (
	TimezoneJakarta = "Asia/Jakarta"
	TimezoneUTC     = "UTC"
)

// NowJakarta mengembalikan time.Time saat ini khusus zona waktu Asia/Jakarta (WIB)
func NowJakarta() *time.Time {
	var result time.Time
	loc, err := time.LoadLocation(TimezoneJakarta)
	if err != nil {
		// Fallback ke FixedZone (+7 jam) jika tzdata sistem tidak ditemukan
		result = time.Now().In(time.FixedZone("WIB", 7*3600))
		return &result
	}

	result = time.Now().In(loc)
	return &result
}

// NowUTC mengembalikan time.Time saat ini dalam UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

func IsNullOrZeroDate(input *time.Time) bool {

	if input == nil || input.IsZero() {
		return true
	}

	return false
}
