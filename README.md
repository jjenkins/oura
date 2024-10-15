# Oura API Go Client

This package provides a Go client for interacting with the Oura API (v2). It allows you to easily retrieve data from your Oura ring, including sleep and readiness information.

## Installation

To use this package in your Go project, you can install it using `go get`:

```bash
go get github.com/jjenkins/ouraring
```

Replace `yourusername` with your actual GitHub username or organization name.

## Usage

Here's a basic example of how to use the Oura API Go client:

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/jjenkins/ouraring"
)

func main() {
    // Create a new client with your Oura API token
    client := ouraring.NewClient("your-api-token")

    // Set the date range for which you want to retrieve data
    startDate := time.Now().AddDate(0, 0, -7) // 7 days ago
    endDate := time.Now()

    // Get sleep data
    sleepData, err := client.GetSleep(startDate, endDate)
    if err != nil {
        log.Fatalf("Error getting sleep data: %v", err)
    }

    // Print sleep data
    for _, sleep := range sleepData {
        fmt.Printf("Sleep on %s: %d minutes\n", sleep.Day, sleep.TotalSleepDuration/60)
    }

    // Get readiness data
    readinessData, err := client.GetReadiness(startDate, endDate)
    if err != nil {
        log.Fatalf("Error getting readiness data: %v", err)
    }

    // Print readiness data
    for _, readiness := range readinessData {
        fmt.Printf("Readiness on %s: Score %d\n", readiness.Day, readiness.Score)
    }
}
```

## Features

- Retrieve sleep data from the Oura API
- Retrieve readiness data from the Oura API
- Easy-to-use client with simple method calls
- Error handling for API requests

## API Documentation

For full details on the Oura API, please refer to the [official Oura API documentation](https://cloud.ouraring.com/docs/).

## Contributing

Contributions to this package are welcome! Here are some ways you can contribute:

1. Report bugs or request features by opening an issue
2. Fix bugs or implement new features by submitting a pull request
3. Improve documentation or add examples

Please ensure that your code adheres to the existing style and that all tests pass before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This is an unofficial Oura API client and is not affiliated with or endorsed by Oura.

## Contact

If you have any questions or feedback, please open an issue on the GitHub repository.

---

Remember to replace `yourusername` with your actual GitHub username or organization name throughout this README.
