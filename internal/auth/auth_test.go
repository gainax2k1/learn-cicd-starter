package auth

// testing GetAPIKey:

import (
	"net/http"
	"testing"
)

// Test functions must take one argument of type *testing.T
// - provides ways to print, skip, and fail test
//    *testing.T

func TestGetAPIKeyEmptyHeader(t *testing.T) {

	header := http.Header{}
	got, err := GetAPIKey(header)
	want := ""
	if got != want {
		t.Errorf("expected string: %v, got: %v", want, got)
	}

	if err == nil || err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

}

func TestGetAPIKeyValidHeader(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKey abc123"},
	}
	got, err := GetAPIKey(header)
	want := "abc123"
	if got != want {
		t.Errorf("expected string: %v, got: %v", want, got)
	}

	if err != nil {
		t.Errorf("expected no error, got error: %v", err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"Bearer abc123"},
	}
	got, err := GetAPIKey(header)
	want := ""
	if got != want {
		t.Errorf("expected string: %v, got: %v", want, got)
	}

	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("Expected err 'malformed authorization header', got: %v", err)
	}
}
