package ouraring

import (
	"fmt"
)

func FormatDuration(seconds int) string {
	hours := seconds / 3600
	remainingSeconds := seconds % 3600
	minutes := remainingSeconds / 60
	secondsRemaining := remainingSeconds % 60

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, secondsRemaining)
}
