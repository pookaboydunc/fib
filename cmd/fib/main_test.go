package main

import (
	"fib/pkg/fibonacci"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestValidQueryParams(t *testing.T) {
	tests := []struct {
		n        int
		big      string
		expected error
	}{
		{-1, "false", ErrNonNegativeNumber},
		{0, "false", nil},
		{1, "false", nil},
		{2, "false", nil},
		{5, "false", nil},
		{7, "false", nil},
		{fibonacci.MAX_32_BIT_N + 1, "false", ErrTooLarge},
		{fibonacci.MAX_64_BIT_N + 1, "false", ErrTooLarge},
		{fibonacci.MAX_64_BIT_N + 1, "true", nil},
	}

	for _, tt := range tests {
		_, err := validQueryParams(&http.Request{URL: &url.URL{RawQuery: fmt.Sprintf("n=%s&big=%s", strconv.Itoa(tt.n), tt.big)}})
		if err != nil {
			if tt.expected == nil {
				t.Errorf("expected no error, got %v", err)
				continue
			}
			if !strings.Contains(err.Error(), tt.expected.Error()) {
				t.Errorf("expected error %v, got %v", tt.expected, err)
			}
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
