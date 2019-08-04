package btpush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	validToken   = "A0AE6EBF-0181-49E7-B7FD-5F8D4B5423C8"
	invalidToken = "F94A3693-16C4-4AF7-A3A5-2B86415C5571"

	validUserUUID   = "CE02E3D2-C16D-4DFC-B36D-82FD3A4F1469"
	invalidUserUUID = "BDF7FA7E-4486-4C1C-807D-D7CC6827B45A"

	sampleDeviceUUID = "CECF9C36-F312-47C1-A5A9-A91E42FDD361"
)

var (
	sampleContent = Content{
		Title: "Sample Title",
		Body:  "Sample Body",
	}
)

func server() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", validToken) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if strings.Contains(string(payload), validUserUUID) {
			fmt.Fprintln(w, `{"status": true}`)
		} else {
			fmt.Fprintln(w, `{"status": false, "error": "oh no!"}`)
		}
	}))
}

func client(url, token string) Client {
	client := New(token)
	client.BaseURL = url
	return client
}

func setup() (*httptest.Server, Client, Client) {
	server := server()
	return server, client(server.URL, validToken), client(server.URL, invalidToken)
}

func Test_Send_Alert_Notifications_User(t *testing.T) {
	server, validClient, invalidClient := setup()
	defer server.Close()

	// Success
	validResp, err := validClient.SendAlertNotificationsUser(validUserUUID, sampleContent)
	if err != nil {
		t.Error(err)
		return
	}
	if !validResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
	}

	invalidResp, err := invalidClient.SendAlertNotificationsUser(invalidUserUUID, sampleContent)
	if err == nil && invalidResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
		return
	}
}

func Test_Send_Alert_Notifications_Devices(t *testing.T) {
	server, validClient, invalidClient := setup()
	defer server.Close()

	// Success
	validResp, err := validClient.SendAlertNotificationsDevices(validUserUUID, []string{sampleDeviceUUID}, sampleContent)
	if err != nil {
		t.Error(err)
		return
	}
	if !validResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
	}

	invalidResp, err := invalidClient.SendAlertNotificationsDevices(validUserUUID, []string{sampleDeviceUUID}, sampleContent)
	if err == nil && invalidResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
		return
	}
}

func Test_Send_Silent_Notifications_User(t *testing.T) {
	server, validClient, invalidClient := setup()
	defer server.Close()

	// Success
	validResp, err := validClient.SendSilentNotificationsUser(validUserUUID, sampleContent)
	if err != nil {
		t.Error(err)
		return
	}
	if !validResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
	}

	invalidResp, err := invalidClient.SendSilentNotificationsUser(invalidUserUUID, sampleContent)
	if err == nil && invalidResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
		return
	}
}

func Test_Send_Silent_Notifications_Devices(t *testing.T) {
	server, validClient, invalidClient := setup()
	defer server.Close()

	// Success
	validResp, err := validClient.SendSilentNotificationsDevices(validUserUUID, []string{sampleDeviceUUID}, sampleContent)
	if err != nil {
		t.Error(err)
		return
	}
	if !validResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
	}

	invalidResp, err := invalidClient.SendSilentNotificationsDevices(invalidUserUUID, []string{sampleDeviceUUID}, sampleContent)
	if err == nil && invalidResp.Status {
		t.Errorf("Error: %s\n", validResp.Error)
		return
	}
}
