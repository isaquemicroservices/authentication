package utils

import "time"

// GetPointerInt64 returns a pointer to the given int64
func GetPointerInt64(value int64) *int64 {
	return &value
}

// GetPointerString returns a pointer to the given string
func GetPointerString(value string) *string {
	return &value
}

// GetPointerFloat64 returns a pointer to the given float64
func GetPointerFloat64(value float64) *float64 {
	return &value
}

// GetPointerTime returns a pointer to the given time
func GetPointerTime(value time.Time) *time.Time {
	return &value
}

// GetPointerBool returns a pointer to the given boolean
func GetPointerBool(value bool) *bool {
	return &value
}
