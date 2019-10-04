package test

import (
	"../handlers"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenURLEmpty(t *testing.T) {
	// Testing the server
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenURL)
	handler.ServeHTTP(rr, req)

	// Check the response body is okay
	expected := "You must enter url parameter\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	println("ShortenURL Endpoint Without Url Parameter Works!")
}

func TestShortenURL(t *testing.T) {

	req, err := http.NewRequest("GET", "?url=www.google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenURL)
	handler.ServeHTTP(rr, req)

	key := rr.Body.String()[1:]
	// Check the response body is okay
	if len(key) != 6 {
		t.Errorf("handler returned unexpected key")
	}

	println("ShortenURL Endpoint Works!")

	req, err = http.NewRequest("GET", key, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(handlers.RedirectURL)
	vars := map[string]string{
		"key": key,
	}
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	expected := `<a href="http://www.google.com">Moved Permanently</a>.` + "\n\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	println("RedirectURL Endpoint Works!")

}
