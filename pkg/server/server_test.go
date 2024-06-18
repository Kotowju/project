// pkg/server/server_test.go

package server_test

import (
	"net/http"
	"net/http/httptest"
	"project/pkg/server"
	"testing"
)

// test jednostkowy serwera
func TestServerStart(t *testing.T) {
	// tworzenie żadania skierowane na ścieżkę /users
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// odpalenie serwera
	go func() {
		err := server.StartServer()
		if err != nil {
			t.Fatalf("StartServer returned an error: %v", err)
		}
	}()

	// opoznienie time.Sleep(1 * time.Second)

	//wykonanie żądania
	http.DefaultClient.Do(req)

	// check statusu
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
