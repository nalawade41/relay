package config

import "errors"

// Config holds every setting the service needs. It is filled once in run.
type Config struct {
	Host         string
	Port         string
	DatabaseURL  string
	MaxBodyBytes int64
}

func LoadConfig(getenv func(string) string) (Config, error) {
	config := Config{
		Host:         getenv("RELAY_HOST"),
		Port:         getenv("RELAY_PORT"),
		DatabaseURL:  getenv("DATABASE_URL"),
		MaxBodyBytes: 64 << 10, // 64 KB
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}

	return config, nil
}
