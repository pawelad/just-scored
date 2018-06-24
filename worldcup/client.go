package worldcup

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// apiURL is the base URL for World Cup API
const apiURL = "http://worldcup.sfg.io/"

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient creates a new World Cup API client
func NewClient() *Client {
	baseURL, _ := url.Parse(apiURL)

	return &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
}

// NewRequest creates an API request to passed endpoint
func (c *Client) NewRequest(method, endpoint string) (*http.Request, error) {
	url, err := c.baseURL.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	return request, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
func (c *Client) Do(request *http.Request, v interface{}) (*http.Response, error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(v)

	return response, err
}
