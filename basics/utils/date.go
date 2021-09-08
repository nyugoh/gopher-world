package utils

import "time"

// Returns the current date time
func CurrentDate() string {
	test("Howdy")
	return time.Now().String()
}
