# bt-push-go-sdk
This is a Go library to interact with the BT Push backend.

### Installation
```shell
go get github.com/brushedtype/bt-push-go-sdk/btpush
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
  btPush := btpush.New(btpush.Config{
    Token: "some-token",
  })

  // Send alert notifications to a user
  resp, err := btPush.SendAlertNotificationsUser("some-uuid", btpush.Content{
    Title:            "Sample Title",
    Body:             "Sample Body",
    Subtitle:         "Sample Subtitle",
    Badge:            99,
    Sound:            "sample.aif",
    ContentAvailable: 1,
  })

  // Error handling
  if err != nil {
    log.Println(err)

    switch err.Type {
      case btpush.ErrorTypeAPNSError:
        log.Println("APNS is experiencing issues")
        break
      default:
        log.Println("Other issues")
        break
    }
  }

  // Send alert notifications to a user's specific devices
  resp, err := btPush.SendAlertNotificationsDevices("some-uuid", []string{"some-device-uuid"}, btpush.Content{
    Title:            "Sample Title",
    Body:             "Sample Body",
    Subtitle:         "Sample Subtitle",
    Badge:            99,
    Sound:            "sample.aif",
    ContentAvailable: 1,
  })
  
  // Send silent notifications to a user's specific devices
  resp, err := btPush.SendSilentNotificationsUser("some-uuid", btpush.Content{
    Data: map[string]string{
      "foo": "bar",
    },
    Badge:          99,
    MutableContent: 1,
    Sound:          "sample.aif",
  })

  // Send silent notifications to a user's specific devices
  resp, err := btPush.SendSilentNotificationsDevices("some-uuid", []string{"some-device-uuid"}, btpush.Content{
    Data: map[string]string{
      "foo": "bar",
    },
    Badge:          99,
    MutableContent: 1,
    Sound:          "sample.aif",
  })
}
```
