package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestStore(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		data := "Hello World!"
		spyStore := &SpyStore{response: data, t: t}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Error, expected %q, got %q", data, response.Body.String())
		}
	})

	t.Run("tells the store to cancel work if request has been cancelled", func(t *testing.T) {
		data := "Hello World!"
		spyStore := &SpyStore{response: data, t: t}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// cancel request
		cancellingCtx, cancelFn := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancelFn)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
