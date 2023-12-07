package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/shadowCow/matchmaker-go/domain/services"
)

var port int = 8888
var address string = fmt.Sprintf("http://localhost:%d", port)
func TestMain(m *testing.M) {
	// setup
	mm := services.NewMatchmakingService()
	a := NewApiHttp(mm)
	go a.Start(port)
	// give the server some time to start
	// a production implementation should do something fancier than sleep
	time.Sleep(100 * time.Millisecond)

	// Run tests
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestFindMatch(t *testing.T) {
	requestBody := `{"playerId":"p1"}`

	// Send a POST request to the test server
	resp := send(t, "find-match", requestBody)
	defer resp.Body.Close()

	wantStatus := http.StatusCreated
	wantBody := `{"playerId":"p1"}
`
	expectResponse(t, resp, wantStatus, wantBody)
}

func TestLeaveMatchmaking(t *testing.T) {
	requestBody := `{"playerId":"p1"}`

	// Send a POST request to the test server
	resp := send(t, "leave-matchmaking", requestBody)
	defer resp.Body.Close()

	wantStatus := http.StatusCreated
	wantBody := `{"playerId":"p1"}
`
	expectResponse(t, resp, wantStatus, wantBody)
}

func send(t *testing.T, endpoint string, body string) *http.Response {
	resp, err := http.Post(
		fmt.Sprintf("%s/%s", address, endpoint),
		"application/json",
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	return resp
}

func expectResponse(t *testing.T, resp *http.Response, wantStatus int, wantBody string) {
// Check the response status code
gotStatus := resp.StatusCode

if gotStatus != wantStatus {
	t.Errorf("Got status code %d, want %d", gotStatus, wantStatus)
}

// Read the response body
responseBody, err := io.ReadAll(resp.Body)
if err != nil {
	t.Fatalf("Error reading response body: %v", err)
}

// Convert the response body to a string
gotBody := string(responseBody)

if gotBody != wantBody {
	t.Errorf("Got body %s, want body %s", gotBody, wantBody)
}
}