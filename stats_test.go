package stats

import (
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStats(t *testing.T) {
	i := 1
	nextCalled := false
	st := StatsHandler{
		Next: httpserver.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			nextCalled = true
			return 0, nil
		}),
	}
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatalf("Test %d: Could not create HTTP request: %v", i, err)
	}

	rec := httptest.NewRecorder()
	st.ServeHTTP(rec, req)

	if !nextCalled {
		t.Errorf("Test %d: Next handler was not called", i)
	}
}

func TestStatsWithValidRedisURL(t *testing.T) {
	i := 1
	nextCalled := false
	st := StatsHandler{
		RedisURL: "redis://localhost:6379",
		Next: httpserver.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			nextCalled = true
			return 0, nil
		}),
	}
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatalf("Test %d: Could not create HTTP request: %v", i, err)
	}

	rec := httptest.NewRecorder()
	st.ServeHTTP(rec, req)

	if !nextCalled {
		t.Errorf("Test %d: Next handler was not called", i)
	}
}
