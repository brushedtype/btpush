package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// Client this is a client to work with the BT Push API
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	token      string
}

// Content represents the data in a notification request
type Content struct {
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Data  interface{} `json:"data"`
}

// Response a response from the server
type Response struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

// AlertUser send alert notifications to a user
func (c *Client) AlertUser(userID string, content Content) (Response, error) {
	return c.POST("/alert", map[string]interface{}{
		"user":    userID,
		"content": content,
	})
}

// AlertDevices send alert notifications to specific devices
func (c *Client) AlertDevices(userID string, devices []string, content Content) (Response, error) {
	return c.POST("/alert", map[string]interface{}{
		"user":    userID,
		"devices": devices,
		"content": content,
	})
}

// POST Make a POST request
func (c *Client) POST(route string, payload interface{}) (Response, error) {
	var r Response
	var url *url.URL
	var err error

	url = c.BaseURL
	url.Path = path.Join(c.BaseURL.Path, route)
	// JSON encoding
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(payload)
	if err != nil {
		return r, err
	}

	// Request creation
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return r, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// Response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	// JSON decoding
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, nil
}

// NewClient create a new API client
func NewClient(token string) *Client {
	url, _ := url.Parse("https://brushedtype-push.herokuapp.com/v0.1")
	return &Client{
		BaseURL:    url,
		httpClient: http.DefaultClient,
		token:      token,
	}
}
