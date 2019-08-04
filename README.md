# bt-push-go-sdk
This is a Go library to interact with the BT Push backend.

### Installation
```shell
go get github.com/brushedtype/bt-push-go-sdk
```

### Example
```go
import (
  "net/http"
  "net/url"
  "log"
  "github.com/brushedtype/bt-push-go-sdk/btpush"
)

func main() {
  // Initialise a client with an auth token
  btPush := btpush.New("some-token")

  // Send alert notifications to a user
  resp, err := btPush.SendAlertNotificationsUser("some-uuid", btpush.Content{
    Title: "Sample Title",
    Body:  "Sample Body",
  })

  // Send alert notifications to a user's specific devices
  resp, err := btPush.SendAlertNotificationsDevices("some-uuid", []string{"some-device-uuid"}, btpush.Content{
    Title: "Sample Title",
    Body:  "Sample Body",
  })
  
  // Send silent notifications to a user's specific devices
  resp, err := btPush.SendSilentNotificationsUser("some-uuid", btpush.Content{
    Title: "Sample Title",
    Body:  "Sample Body",
  })

  // Send silent notifications to a user's specific devices
  resp, err := btPush.SendSilentNotificationsDevices("some-uuid", []string{"some-device-uuid"}, btpush.Content{
    Title: "Sample Title",
    Body:  "Sample Body",
  })
}
```
