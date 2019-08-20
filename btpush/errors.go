package btpush

// ErrorType error type
type ErrorType string

const (
	ErrorTypeRequestBadPayload       ErrorType = "request_bad_payload"
	ErrorTypeRequestInvalidUser      ErrorType = "request_invalid_user"
	ErrorTypeRequestInvalidDevice    ErrorType = "request_invalid_device"
	ErrorTypeRequestUnauthorized     ErrorType = "request_unauthorized"
	ErrorTypeRequestMissingAuthToken ErrorType = "request_missing_auth_token"

	ErrorTypeServerUserNoDevices ErrorType = "server_user_no_devices"
	ErrorTypeServerAuthKey       ErrorType = "server_auth_key_error"
	ErrorTypeServerError         ErrorType = "server_error" // generic server error

	ErrorTypeAPNSError ErrorType = "apns_error" // generic APNS error
	ErrorTypeOther     ErrorType = "other"
)

// Error represents an error in a HTTP request/response
type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}
