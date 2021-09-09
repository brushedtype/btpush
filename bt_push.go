package btpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	baseURL = "https://brushedtype-push.herokuapp.com/v0.1"
)

// Client this is a client to work with the BT Push API
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Config     Config
	Debug      bool
}

// Config Represents a Client configuration
type Config struct {
	Token string
}

const (
	TimeSensitiveInturruptionLevel = "time-sensitive"
	ActiveInterruptionLevel        = "active"
	PassiveInterruptionLevel       = "passive"
)

// Content represents the data in a notification request
type Content struct {
	Data              interface{} `json:"data,omitempty"`
	ContentAvailable  int         `json:"content_available,omitempty"`
	InterruptionLevel string      `json:"interruption-level,omitempty"`
	Badge             int         `json:"badge,omitempty"`
	Category          string      `json:"category,omitempty"`
	MutableContent    int         `json:"mutable_content,omitempty"`
	ThreadID          string      `json:"thread_id,omitempty"`
	Sound             interface{} `json:"sound,omitempty"`
	Expiration        time.Time   `json:"expiration,omitempty"`

	// Alert
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	TitleLocArgs []string `json:"title_loc_args,omitempty"`
	TitleLocKey  string   `json:"title_loc_key,omitempty"`
	LocArgs      []string `json:"loc_args,omitempty"`
	LocKey       string   `json:"loc_key,omitempty"`
}

// Response a response from the server
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   *Error `json:"error"`
}

// ClientError represents an error in the client or response
type ClientError struct {
	message string
	Type    ErrorType
}

func (r ClientError) Error() string {
	return r.message
}

// SendAlertNotificationsUser send alert notifications to a user
func (c *Client) SendAlertNotificationsUser(userID string, content Content) (Response, *ClientError) {
	return c.POST("/alert", map[string]interface{}{
		"user":    userID,
		"content": content,
	})
}

// SendAlertNotificationsDevices send alert notifications to specific devices
func (c *Client) SendAlertNotificationsDevices(userID string, devices []string, content Content) (Response, *ClientError) {
	return c.POST("/alert", map[string]interface{}{
		"user":    userID,
		"devices": devices,
		"content": content,
	})
}

// SendSilentNotificationsUser send silent notifications to a user
func (c *Client) SendSilentNotificationsUser(userID string, content Content) (Response, *ClientError) {
	return c.POST("/silent", map[string]interface{}{
		"user":    userID,
		"content": content,
	})
}

// SendSilentNotificationsDevices send silent notifications to specific devices
func (c *Client) SendSilentNotificationsDevices(userID string, devices []string, content Content) (Response, *ClientError) {
	return c.POST("/silent", map[string]interface{}{
		"user":    userID,
		"devices": devices,
		"content": content,
	})
}

// POST Make a POST request
func (c *Client) POST(route string, payload interface{}) (Response, *ClientError) {
	var r Response
	var url *url.URL
	var err error

	url, _ = url.Parse(c.BaseURL)
	url.Path = path.Join(url.Path, route)
	// JSON encoding
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(payload)
	if err != nil {
		return r, &ClientError{
			Type:    ErrorTypeOther,
			message: err.Error(),
		}
	}

	if c.Debug {
		log.Printf("POST '%s' with %d-byte payload: %+v\n", url.String(), len(jsonBytes), string(jsonBytes))
	}

	// Request creation
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return r, &ClientError{
			Type:    ErrorTypeOther,
			message: err.Error(),
		}
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.Token))

	// Response
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return r, &ClientError{
			Type:    ErrorTypeOther,
			message: err.Error(),
		}
	}
	defer resp.Body.Close()

	if c.Debug {
		log.Printf("HTTP %s\n", resp.Status)
	}

	// JSON decoding
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return r, &ClientError{
			Type:    ErrorTypeOther,
			message: err.Error(),
		}
	} else if r.Error != nil {
		return r, &ClientError{
			Type:    r.Error.Type,
			message: r.Error.Message,
		}
	}

	return r, nil
}

// New create an API client with usual defaults
func New(config Config) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Config:     config,
		Debug:      false,
	}
}

// NewDebug create an API client with usual defaults and debugging turned on
func NewDebug(config Config) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Config:     config,
		Debug:      true,
	}
}
