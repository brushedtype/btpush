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
  client := sdk.Client{
    HTTPClient: http.DefaultClient,
    Token:      "some-token",
  }
  resp, err := client.SendSilentNotificationsUser("some-uuid", sdk.Content{
    Title: "Sample Title",
    Body:  "Sample Body",
  })
  if err != nil {
    log.Println(err)
    return
  }
  if !resp.Status {
    log.Printf("Error: %s\n", resp.Error)
  }
}
```
