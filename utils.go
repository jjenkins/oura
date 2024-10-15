// Package ouraring provides a client for interacting with the Oura Ring API.
package ouraring

import (
	"fmt"
)

// FormatDuration converts a duration in seconds to a formatted string.
//
// The function takes a duration in seconds and returns a string in the format "XXh YYm ZZs",
// where XX is hours, YY is minutes, and ZZ is seconds. Each component is zero-padded to two digits.
//
// Parameters:
//   - seconds: The duration in seconds to be formatted.
//
// Returns:
//   - A string representing the duration in the format "XXh YYm ZZs".
//
// Example:
//
//	FormatDuration(3661) returns "01h 01m 01s"
//	FormatDuration(7200) returns "02h 00m 00s"
func FormatDuration(seconds int) string {
	hours := seconds / 3600
	remainingSeconds := seconds % 3600
	minutes := remainingSeconds / 60
	secondsRemaining := remainingSeconds % 60

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, secondsRemaining)
}
