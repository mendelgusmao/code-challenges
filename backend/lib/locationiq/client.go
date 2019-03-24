package locationiq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const locationIQURL = "https://us1.locationiq.com/v1/search.php"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) Geocode(address string) (float64, float64, error) {
	parameters := url.Values{}
	parameters.Add("key", c.apiKey)
	parameters.Add("countrycodes", "br")
	parameters.Add("q", address)
	parameters.Add("format", "json")

	uri := fmt.Sprintf("%s?%s", locationIQURL, parameters.Encode())
	res, err := http.Get(uri)

	if err != nil {
		return 0.0, 0.0, fmt.Errorf("geocoder: %s", err)
	}

	response := []map[string]interface{}{}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return 0.0, 0.0, fmt.Errorf("geocoder: %s", err)
	}

	if len(response) == 0 {
		return 0.0, 0.0, fmt.Errorf("geocoder: address not found")
	}

	lat, _ := strconv.ParseFloat(response[0]["lat"].(string), 64)
	long, _ := strconv.ParseFloat(response[0]["lon"].(string), 64)

	return lat, long, nil
}
