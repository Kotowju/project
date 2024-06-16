// pkg/server/server_test.go

package server_test

import (
	"net/http"
	"net/http/httptest"
	"project/pkg/server"
	"testing"
)

// TestServerStart function tests the StartServer function.
func TestServerStart(t *testing.T) {
	// Create a new HTTP request to test the server
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the StartServer function (assuming it starts the server internally)
	go func() {
		err := server.StartServer()
		if err != nil {
			t.Fatalf("StartServer returned an error: %v", err)
		}
	}()

	// Wait for the server to start (optional, depending on your setup)
	// time.Sleep(1 * time.Second)

	// Perform the HTTP request against the server
	http.DefaultClient.Do(req)

	// Check the status code (assuming the handler responds with 201 Created)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Additional assertions if needed
}
