package config

import (
	"os"
	"time"

	"github.com/evalphobia/ipqualityscore-go/client"
)

const (
	envAPIKey = "IPQS_APIKEY"
)

// Config contains parameters for IPQualityScore.
type Config struct {
	APIKey string

	Debug   bool
	Timeout time.Duration
}

// Client generates client.Client from Config data.
func (c Config) Client() (*client.Client, error) {
	cli := client.New()
	cli.SetOption(client.Option{
		Debug:   c.Debug,
		Timeout: c.Timeout,
	})
	cli.SetAPIKey(c.getAPIKey())
	return cli, nil
}

// getAPIKey returns API Key from Config data or Environment variables.
func (c Config) getAPIKey() string {
	apiKey := os.Getenv(envAPIKey)
	if apiKey != "" {
		return apiKey
	}
	return c.APIKey
}
