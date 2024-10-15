// Package ouraring provides a client for interacting with the Oura Ring API.
package ouraring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetReadiness retrieves the daily readiness data for a specified date range.
//
// It sends a GET request to the Oura API's daily_readiness endpoint and returns
// a slice of Readiness structs containing the data for each day in the range.
//
// Parameters:
//   - startDate: The start date of the range (inclusive).
//   - endDate: The end date of the range (inclusive).
//
// Returns:
//   - A slice of Readiness structs containing the daily readiness data.
//   - An error if the API request fails or the response cannot be parsed.
func (c *Client) GetReadiness(startDate, endDate time.Time) ([]Readiness, error) {
	// Construct the URL with query parameters for the date range
	url := fmt.Sprintf("%s/usercollection/daily_readiness?start_date=%s&end_date=%s",
		c.baseURL, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the Authorization header with the client's token
	req.Header.Set("Authorization", "Bearer "+c.token)

	// Send the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is not OK (200)
	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	// Define a struct to hold the response data
	var response struct {
		Data []Readiness `json:"data"`
	}

	// Decode the JSON response body into the response struct
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	// Return the slice of Readiness data
	return response.Data, nil
}
