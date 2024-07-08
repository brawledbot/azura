package azura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Client represents the main API client for interacting with Azura.
type Client struct {
	BaseURL    string       // BaseURL is the base URL for the Azura API.
	HTTPClient *http.Client // HTTPClient is the HTTP client used to make requests.
	APIKey     string       // APIKey is the API key used for authentication.
}

// Response represents a response from the Azura API.
type Response struct {
	Code    int         `json:"code"`    // Code contains the status code of the response.
	Data    interface{} `json:"data"`    // Data contains the returned data from the API.
	Message string      `json:"message"` // Message contains any informative message returned by the API.
}

// NewClient creates a new instance of the Azura API client.
func NewClient(baseURL string, apiKey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
	}
}

// DoRequest makes an HTTP request to the Azura API and returns the API response or an error.
// method: HTTP method (GET, POST, PUT, DELETE, etc.)
// endpoint: API endpoint relative to the base URL.
// body: Optional. Body data to send with the request (typically for POST and PUT requests).
func (c *Client) DoRequest(method, endpoint string, body interface{}) (*Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	var req *http.Request
	var err error

	if body != nil {
		var jsonBody []byte
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			return nil, fmt.Errorf("the azura client cannot connect to the API server")
		}

		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResponse Response
	err = json.Unmarshal(responseData, &jsonResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, jsonResponse.Message)
	}

	return &jsonResponse, nil
}
