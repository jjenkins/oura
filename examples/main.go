package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jjenkins/ouraring"
)

func main() {
	client := ouraring.NewClient(os.Getenv("OURA_ACCESS_TOKEN"))

	startDate := time.Now().AddDate(0, 0, -7) // 7 days ago
	endDate := time.Now()

	// sleepData, err := client.GetSleep(startDate, endDate)
	// if err != nil {
	// 	log.Fatalf("Error getting sleep data: %v", err)
	// }

	readinessData, err := client.GetReadiness(startDate, endDate)
	if err != nil {
		log.Fatalf("Error getting readiness data: %v", err)
	}

	for _, readiness := range readinessData {
		fmt.Printf("Temperature %d\n", readiness.Contributors.BodyTemperature)
		fmt.Printf("Temperature Deviation %.2f\n", readiness.TemperatureDeviation)
		fmt.Printf("Temperature Deviation Trend %.2f\n", readiness.TemperatureTrendDeviation)
		fmt.Printf("Readiness on %s: Score %d\n\n", readiness.Day, readiness.Score)
	}

	// for _, sleep := range sleepData {
	// 	fmt.Printf("ID %s\n", sleep.ID)
	// 	fmt.Printf("Type %s\n", sleep.Type)
	// 	fmt.Printf("Efficiency %d\n", sleep.Efficiency)
	// 	fmt.Printf("Average Breath %.2f\n", sleep.AverageBreath)
	// 	fmt.Printf("Average Heart Rate %.2f\n", sleep.AverageHeartRate)
	// 	fmt.Printf("Lowest Heart Rate %d\n", sleep.LowestHeartRate)
	// 	fmt.Printf("Average HRV %d\n", sleep.AverageHRV)
	// 	fmt.Printf("Period %d\n", sleep.Period)
	// 	fmt.Printf("Restless Periods %d\n", sleep.RestlessPeriods)
	// 	fmt.Printf("REM Sleep Duration %s\n", FormatDuration(sleep.REMSleepDuration))
	// 	fmt.Printf("Deep Sleep Duration %s\n", FormatDuration(sleep.DeepSleepDuration))
	// 	fmt.Printf("Light Sleep Duration %s\n", FormatDuration(sleep.LightSleepDuration))
	// 	fmt.Printf("Awake Time %s\n", FormatDuration(sleep.AwakeTime))
	// 	fmt.Printf("Deep Sleep Duration %s\n", FormatDuration(sleep.DeepSleepDuration))
	// 	fmt.Printf("Latency %s\n", FormatDuration(sleep.Latency))
	// 	fmt.Printf("Sleep on %s: %s\n", sleep.Day, FormatDuration(sleep.TotalSleepDuration))
	// 	fmt.Printf("Time in bed %s: %s\n\n", sleep.Day, FormatDuration(sleep.TimeInBed))
	// }
}
