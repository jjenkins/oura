package ouraring

import "time"

type Sleep struct {
	ID                 string    `json:"id"`
	AverageBreath      float64   `json:"average_breath"`
	AverageHeartRate   float64   `json:"average_heart_rate"`
	AverageHRV         int       `json:"average_hrv"`
	AwakeTime          int       `json:"awake_time"`
	BedtimeEnd         time.Time `json:"bedtime_end"`
	BedtimeStart       time.Time `json:"bedtime_start"`
	Day                string    `json:"day"`
	DeepSleepDuration  int       `json:"deep_sleep_duration"`
	Efficiency         int       `json:"efficiency"`
	Latency            int       `json:"latency"`
	LightSleepDuration int       `json:"light_sleep_duration"`
	LowBatteryAlert    bool      `json:"low_battery_alert"`
	LowestHeartRate    int       `json:"lowest_heart_rate"`
	Period             int       `json:"period"`
	REMSleepDuration   int       `json:"rem_sleep_duration"`
	RestlessPeriods    int       `json:"restless_periods"`
	TimeInBed          int       `json:"time_in_bed"`
	TotalSleepDuration int       `json:"total_sleep_duration"`
	Type               string    `json:"type"`
}

type Readiness struct {
	ID                        string       `json:"id"`
	Contributors              Contributors `json:"contributors"`
	Day                       string       `json:"day"`
	Score                     int          `json:"score"`
	TemperatureDeviation      float64      `json:"temperature_deviation"`
	TemperatureTrendDeviation float64      `json:"temperature_trend_deviation"`
	Timestamp                 time.Time    `json:"timestamp"`
}

type Contributors struct {
	ActivityBalance     int `json:"activity_balance"`
	BodyTemperature     int `json:"body_temperature"`
	HRVBalance          int `json:"hrv_balance"`
	PreviousDayActivity int `json:"previous_day_activity"`
	PreviousNight       int `json:"previous_night"`
	RecoveryIndex       int `json:"recovery_index"`
	RestingHeartRate    int `json:"resting_heart_rate"`
	SleepBalance        int `json:"sleep_balance"`
}
