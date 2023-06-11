package internals_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestRunServer(t *testing.T) {
	os.Setenv("AUTH_USERNAME", "testuser")
	os.Setenv("AUTH_PASSWORD", "testpass")
	defer os.Unsetenv("AUTH_USERNAME")
	defer os.Unsetenv("AUTH_PASSWORD")

	go internals.RunServer()

	time.Sleep(time.Second)

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("testuser", "testpass")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
