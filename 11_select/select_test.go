package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Selects the fast server appropriately", func(t *testing.T) {
		// Create two http test servers, one fast and one slow
		slowServer := createDelayedServer(20 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got, err := Racer(slowUrl, fastUrl, 1*time.Second)
		want := fastUrl

		if err != nil {
			t.Errorf("Unexpected got an error when was not expecting one")
		}

		if got != want {
			t.Errorf("Unexpected race result, got %q, want %q", got, want)
		}
	})

	t.Run("Errors when waiting for longer than x seconds", func(t *testing.T) {
		slowServer := createDelayedServer(200 * time.Millisecond)
		defer slowServer.Close()

		_, err := Racer(slowServer.URL, slowServer.URL, 100*time.Millisecond)

		if err == nil {
			t.Errorf("Expected an error but did not get one")
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
