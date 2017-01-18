// handlers_test.go
package handlers_test

import (
	"github.com/steventux/obd2-data-api/handlers"
	"github.com/tborisova/clean_like_gopher"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateObd2DataHandler(t *testing.T) {
	m, _ := clean_like_gopher.NewCleaningGopher("mongo",
		map[string]string{"host": "localhost", "dbName": "test", "port": "27017"})

	// Create a request to pass to our handler.
	req, _ := http.NewRequest("GET", "/", nil)

	query := req.URL.Query()
	query.Add("k42", "14.4")
	query.Add("kff1202", "120")
	query.Add("kff1225", "198")

	req.URL.RawQuery = query.Encode()

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateObd2Data)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	m.Clean(nil)
	m.Close()
}

func TestCreateObd2DataHandlerAuthentication(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	query := req.URL.Query()
	query.Add("id", "123abc")
	query.Add("kff1202", "120")
	query.Add("kff1225", "198")

	req.URL.RawQuery = query.Encode()

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateObd2Data)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 403 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 403)
	}

	// Check the response body is what we expect.
	expected := "Forbidden"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
