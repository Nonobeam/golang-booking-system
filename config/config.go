package config

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		PG      PG
		RMQ     RMQ
		Metrics Metrics
		Swagger Swagger
	}

	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.App.Name = getEnvRequired("APP_NAME")
	cfg.App.Version = getEnvRequired("APP_VERSION")

	cfg.HTTP.Port = getEnvRequired("HTTP_PORT")
	cfg.HTTP.UsePreforkMode = getEnvBool("HTTP_USE_PREFORK_MODE", false)

	cfg.Log.Level = getEnvRequired("LOG_LEVEL")

	cfg.PG.URL = getEnvRequired("PG_URL")
	poolMax, err := strconv.Atoi(getEnvRequired("PG_POOL_MAX"))
	if err != nil {
		return nil, fmt.Errorf("invalid PG_POOL_MAX: %w", err)
	}
	cfg.PG.PoolMax = poolMax

	cfg.RMQ.ServerExchange = getEnvRequired("RMQ_RPC_SERVER")
	cfg.RMQ.ClientExchange = getEnvRequired("RMQ_RPC_CLIENT")
	cfg.RMQ.URL = getEnvRequired("RMQ_URL")

	cfg.Metrics.Enabled = getEnvBool("METRICS_ENABLED", true)

	cfg.Swagger.Enabled = getEnvBool("SWAGGER_ENABLED", false)

	return cfg, nil
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
