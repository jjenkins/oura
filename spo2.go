package ouraring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetDailySpO2 retrieves the daily SpO2 data for a specified date range.
//
// It sends a GET request to the Oura API's daily_spo2 endpoint and returns
// a slice of DailySpO2Model structs containing the data for each day in the range.
//
// Parameters:
//   - startDate: The start date of the range (inclusive).
//   - endDate: The end date of the range (inclusive).
//
// Returns:
//   - A slice of DailySpO2Model structs containing the daily SpO2 data.
//   - An error if the API request fails or the response cannot be parsed.
func (c *Client) GetDailySpO2(startDate, endDate time.Time) ([]DailySpO2Model, error) {
	// Construct the URL with query parameters for the date range
	url := fmt.Sprintf("%s/usercollection/daily_spo2?start_date=%s&end_date=%s",
		c.baseURL, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Authorization header with the client's token
	req.Header.Set("Authorization", "Bearer "+c.token)

	// Send the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is not OK (200)
	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	// Define a struct to hold the response data
	var response struct {
		Data []DailySpO2Model `json:"data"`
	}

	// Decode the JSON response body into the response struct
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return response.Data, nil
}

// GetSingleDaySpO2 retrieves the SpO2 data for a specific day using its ID.
//
// It sends a GET request to the Oura API's daily_spo2 endpoint for a single day
// and returns a DailySpO2Model struct containing the data for that day.
//
// Parameters:
//   - id: The unique identifier for the specific day's SpO2 data.
//
// Returns:
//   - A pointer to a DailySpO2Model struct containing the daily SpO2 data.
//   - An error if the API request fails or the response cannot be parsed.
func (c *Client) GetSingleDaySpO2(id string) (*DailySpO2Model, error) {
	// Construct the URL for the specific SpO2 data
	url := fmt.Sprintf("%s/usercollection/daily_spo2/%s", c.baseURL, id)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Authorization header with the client's token
	req.Header.Set("Authorization", "Bearer "+c.token)

	// Send the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is not OK (200)
	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	// Define a variable to hold the response data
	var spo2Data DailySpO2Model

	// Decode the JSON response body into the spo2Data struct
	err = json.NewDecoder(resp.Body).Decode(&spo2Data)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &spo2Data, nil
}
