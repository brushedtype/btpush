package sdk

import "net/http"

// Client this is a client to work with the BT Push API
type Client struct {
	BaseURL    string
	httpClient *http.Client
}

// Content represents the data in a notification request
type Content struct {
	Title string                      `json:"title"`
	Body  string                      `json:"body"`
	Data  map[interface{}]interface{} `json:"data"`
}

// AlertUser send alert notifications to a user
func (c *Client) AlertUser(userID string, content Content) {

}

// AlertDevices send alert notifications to specific devices
func (c *Client) AlertDevices(userID string, devices []string, content Content) {

}
