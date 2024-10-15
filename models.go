// Package ouraring provides a client for the Oura Ring API.
package ouraring

import "time"

// Sleep represents a sleep session recorded by the Oura Ring.
type Sleep struct {
	ID                 string    `json:"id"`                   // Unique identifier for the sleep session
	AverageBreath      float64   `json:"average_breath"`       // Average breathing rate during sleep (breaths per minute)
	AverageHeartRate   float64   `json:"average_heart_rate"`   // Average heart rate during sleep (beats per minute)
	AverageHRV         int       `json:"average_hrv"`          // Average heart rate variability during sleep (milliseconds)
	AwakeTime          int       `json:"awake_time"`           // Time spent awake during the sleep session (seconds)
	BedtimeEnd         time.Time `json:"bedtime_end"`          // Time when the sleep session ended
	BedtimeStart       time.Time `json:"bedtime_start"`        // Time when the sleep session started
	Day                string    `json:"day"`                  // Date of the sleep session (YYYY-MM-DD)
	DeepSleepDuration  int       `json:"deep_sleep_duration"`  // Time spent in deep sleep (seconds)
	Efficiency         int       `json:"efficiency"`           // Sleep efficiency (0-100)
	Latency            int       `json:"latency"`              // Time taken to fall asleep (seconds)
	LightSleepDuration int       `json:"light_sleep_duration"` // Time spent in light sleep (seconds)
	LowBatteryAlert    bool      `json:"low_battery_alert"`    // Whether a low battery alert occurred during the session
	LowestHeartRate    int       `json:"lowest_heart_rate"`    // Lowest heart rate recorded during sleep (beats per minute)
	Period             int       `json:"period"`               // Sleep period identifier
	REMSleepDuration   int       `json:"rem_sleep_duration"`   // Time spent in REM sleep (seconds)
	RestlessPeriods    int       `json:"restless_periods"`     // Number of restless periods during sleep
	TimeInBed          int       `json:"time_in_bed"`          // Total time spent in bed (seconds)
	TotalSleepDuration int       `json:"total_sleep_duration"` // Total time spent sleeping (seconds)
	Type               string    `json:"type"`                 // Type of sleep session (e.g., "long_sleep", "short_sleep")
}

// Readiness represents the daily readiness score and its components.
type Readiness struct {
	ID                        string       `json:"id"`                          // Unique identifier for the readiness record
	Contributors              Contributors `json:"contributors"`                // Breakdown of factors contributing to the readiness score
	Day                       string       `json:"day"`                         // Date of the readiness score (YYYY-MM-DD)
	Score                     int          `json:"score"`                       // Overall readiness score (0-100)
	TemperatureDeviation      float64      `json:"temperature_deviation"`       // Deviation from the user's average body temperature
	TemperatureTrendDeviation float64      `json:"temperature_trend_deviation"` // Deviation from the user's temperature trend
	Timestamp                 time.Time    `json:"timestamp"`                   // Time when the readiness score was calculated
}

// Contributors represents the various factors that contribute to the readiness score.
type Contributors struct {
	ActivityBalance     int `json:"activity_balance"`      // Contribution from activity balance (0-100)
	BodyTemperature     int `json:"body_temperature"`      // Contribution from body temperature (0-100)
	HRVBalance          int `json:"hrv_balance"`           // Contribution from heart rate variability balance (0-100)
	PreviousDayActivity int `json:"previous_day_activity"` // Contribution from previous day's activity (0-100)
	PreviousNight       int `json:"previous_night"`        // Contribution from previous night's sleep (0-100)
	RecoveryIndex       int `json:"recovery_index"`        // Contribution from recovery index (0-100)
	RestingHeartRate    int `json:"resting_heart_rate"`    // Contribution from resting heart rate (0-100)
	SleepBalance        int `json:"sleep_balance"`         // Contribution from sleep balance (0-100)
}
