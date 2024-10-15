package ouraring

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) GetReadiness(startDate, endDate time.Time) ([]Readiness, error) {
	url := fmt.Sprintf("%s/usercollection/daily_readiness?start_date=%s&end_date=%s",
		c.baseURL, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, NewAPIError(resp)
	}

	var response struct {
		Data []Readiness `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
