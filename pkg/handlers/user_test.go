package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"project/pkg/handlers"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	// fałszywe żądanie z danymi użytkownika + nowe żądanie HTTP typu POST na endpoint /users
	requestBody := []byte(`{"username": "testuser", "email": "testuser@example.com"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateUser)

	handler.ServeHTTP(rr, req)

	// sprawdza kod odpowiedzi;czy jest 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}
func TestCreateKeyHandler(t *testing.T) {
	//analogicznie
	requestBody := []byte(`{"name": "testkey", "secret": "testsecret", "username": "testuser"}`)
	req, err := http.NewRequest("POST", "/keys", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateKey)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func TestCreateItemHandler(t *testing.T) {

	requestBody := []byte(`{"sector": "testsector", "name": "testitem", "description": "testdescription"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateItem)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func TestCreatePolicyHandler(t *testing.T) {
	requestBody := []byte(`{"name": "testpolicy", "definition": "testdefinition"}`)
	req, err := http.NewRequest("POST", "/policies", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreatePolicy)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}
