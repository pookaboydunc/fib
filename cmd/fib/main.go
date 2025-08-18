package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"fib/pkg/fibonacci"
)

var (
	ErrMissingQueryParam = errors.New("missing query parameter 'n'")
	ErrInvalidNumber     = errors.New("invalid number")
	ErrNonNegativeNumber = errors.New("n must be non-negative")
	ErrTooLarge          = errors.New("n too large")
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
}

func validQueryParams(r *http.Request) (int, error) {
	// check for missing 'n' query param
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		return 0, ErrMissingQueryParam
	}

	// check if 'n' is a valid number
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, ErrInvalidNumber
	}

	// if 'n' is larger than the safe limit check if the 'big' query param is set
	if limit := fibonacci.SafeLimit(); n > limit {
		big := r.URL.Query().Get("big")
		if big != "true" {
			return n, fmt.Errorf("%s; max supported is %d in 64 bit environments and %d in 32 bit environments", ErrTooLarge.Error(), fibonacci.MAX_64_BIT_N, fibonacci.MAX_32_BIT_N)
		}
	}
	return n, nil
}

func nthFibSequence(w http.ResponseWriter, r *http.Request) {
	// validate query params
	n, err := validQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// request fibonacci sequence up to n numbers long
	sequence, err := fibonacci.Sequence(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sequence)
}

func main() {
	// Parse command line flags
	configPath := flag.String("config", "config/config.json", "Path to configuration file")
	flag.Parse()

	// Load configuration from file
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	http.HandleFunc("/fibonacci", nthFibSequence)
	log.Printf("Fibonacci service running on http://%s/fibonacci", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func loadConfig(path string) (*Config, error) {
	// Ensure config directory exists
	configDir := filepath.Dir(path)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	// Check if config file exists, create default if not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		defaultConfig := Config{}
		defaultConfig.Server.Host = "0.0.0.0"
		defaultConfig.Server.Port = 8080

		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %v", err)
		}

		if err := os.WriteFile(path, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config: %v", err)
		}

		return &defaultConfig, nil
	}

	// Read and parse config file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
