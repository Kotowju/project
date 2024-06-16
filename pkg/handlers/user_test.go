package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"project/pkg/handlers"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	// Mock request body
	requestBody := []byte(`{"username": "testuser", "email": "testuser@example.com"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function from the handler method
	handler := http.HandlerFunc(handlers.CreateUser)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Additional assertions if needed
}
func TestCreateKeyHandler(t *testing.T) {
	// Mock request body
	requestBody := []byte(`{"name": "testkey", "secret": "testsecret", "username": "testuser"}`)
	req, err := http.NewRequest("POST", "/keys", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function from the handler method
	handler := http.HandlerFunc(handlers.CreateKey)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Additional assertions if needed
}

func TestCreateItemHandler(t *testing.T) {
	// Mock request body
	requestBody := []byte(`{"sector": "testsector", "name": "testitem", "description": "testdescription"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function from the handler method
	handler := http.HandlerFunc(handlers.CreateItem)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Additional assertions if needed
}

func TestCreatePolicyHandler(t *testing.T) {
	// Mock request body
	requestBody := []byte(`{"name": "testpolicy", "definition": "testdefinition"}`)
	req, err := http.NewRequest("POST", "/policies", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function from the handler method
	handler := http.HandlerFunc(handlers.CreatePolicy)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Additional assertions if needed
}
