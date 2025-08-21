package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestValidQueryParams(t *testing.T) {
	tests := []struct {
		n        int
		expected error
	}{
		{-1, ErrNonNegativeNumber},
		{0, nil},
		{1, nil},
		{2, nil},
		{5, nil},
		{7, nil},
	}

	for _, tt := range tests {
		_, err := validQueryParams(&http.Request{URL: &url.URL{RawQuery: fmt.Sprintf("n=%s", strconv.Itoa(tt.n))}})
		if err != tt.expected {
			t.Errorf("expected error %v, got %v", tt.expected, err)
		}
	}

	//missing query param
	_, err := validQueryParams(&http.Request{URL: &url.URL{RawQuery: ""}})
	if err != ErrMissingQueryParam {
		t.Errorf("expected error %v, got %v", ErrMissingQueryParam, err)
	}

	//invalid number query param
	_, err = validQueryParams(&http.Request{URL: &url.URL{RawQuery: "n=abc"}})
	if err != ErrInvalidNumber {
		t.Errorf("expected error %v, got %v", ErrInvalidNumber, err)
	}
}
