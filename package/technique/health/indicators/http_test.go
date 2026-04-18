package indicators

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPIndicator_Up(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ind := NewHTTPIndicator("payment-api", srv.URL, srv.Client())
	result := ind.Check(context.Background())

	if !result.Up {
		t.Errorf("Up = false, want true; error = %q", result.Error)
	}
	if result.Details["statusCode"] != http.StatusOK {
		t.Errorf("Details.statusCode = %v, want %d", result.Details["statusCode"], http.StatusOK)
	}
}

func TestHTTPIndicator_Down_Non2xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	ind := NewHTTPIndicator("payment-api", srv.URL, srv.Client())
	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false for 503")
	}
	if result.Error == "" {
		t.Error("Error should describe the unexpected status")
	}
	if result.Details["statusCode"] != http.StatusServiceUnavailable {
		t.Errorf("Details.statusCode = %v, want %d", result.Details["statusCode"], http.StatusServiceUnavailable)
	}
}

func TestHTTPIndicator_Down_Unreachable(t *testing.T) {
	ind := NewHTTPIndicator("dead-service", "http://127.0.0.1:1", nil)
	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false for unreachable host")
	}
	if result.Error == "" {
		t.Error("Error should be set when request fails")
	}
}

func TestHTTPIndicator_Down_InvalidURL(t *testing.T) {
	ind := NewHTTPIndicator("bad-url", "://invalid", nil)
	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false for invalid URL")
	}
}

func TestHTTPIndicator_DefaultClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	// nil client → should use http.DefaultClient
	ind := NewHTTPIndicator("svc", srv.URL, nil)
	if ind.client != http.DefaultClient {
		t.Error("nil client should fall back to http.DefaultClient")
	}
}

func TestHTTPIndicator_2xxVariants(t *testing.T) {
	cases := []int{200, 201, 204}
	for _, code := range cases {
		code := code
		t.Run(http.StatusText(code), func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(code)
			}))
			defer srv.Close()

			ind := NewHTTPIndicator("svc", srv.URL, srv.Client())
			result := ind.Check(context.Background())

			if !result.Up {
				t.Errorf("status %d should be Up", code)
			}
		})
	}
}
