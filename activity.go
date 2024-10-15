package ouraring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetDailyActivity retrieves the daily activity data for a specified date range.
//
// It sends a GET request to the Oura API's daily_activity endpoint and returns
// a slice of DailyActivityModel structs containing the data for each day in the range.
//
// Parameters:
//   - startDate: The start date of the range (inclusive).
//   - endDate: The end date of the range (inclusive).
//
// Returns:
//   - A slice of DailyActivityModel structs containing the daily activity data.
//   - An error if the API request fails or the response cannot be parsed.
func (c *Client) GetDailyActivity(startDate, endDate time.Time) ([]DailyActivity, error) {
	url := fmt.Sprintf("%s/usercollection/daily_activity?start_date=%s&end_date=%s",
		c.baseURL, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	var response struct {
		Data []DailyActivity `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return response.Data, nil
}

// GetSingleDayActivity retrieves the activity data for a specific day using its ID.
//
// It sends a GET request to the Oura API's daily_activity endpoint for a single day
// and returns a DailyActivityModel struct containing the data for that day.
//
// Parameters:
//   - id: The unique identifier for the specific day's activity data.
//
// Returns:
//   - A pointer to a DailyActivityModel struct containing the daily activity data.
//   - An error if the API request fails or the response cannot be parsed.
func (c *Client) GetSingleDayActivity(id string) (*DailyActivity, error) {
	url := fmt.Sprintf("%s/usercollection/daily_activity/%s", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	var activityData DailyActivity

	err = json.NewDecoder(resp.Body).Decode(&activityData)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &activityData, nil
}
