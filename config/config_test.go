package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("APP_NAME", "test-app")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("PG_URL", "postgres://test")
	os.Setenv("PG_POOL_MAX", "10")
	os.Setenv("RMQ_RPC_SERVER", "test-server")
	os.Setenv("RMQ_RPC_CLIENT", "test-client")
	os.Setenv("RMQ_URL", "amqp://test")

	cfg, err := NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() failed: %v", err)
	}

	if cfg.App.Name != "test-app" {
		t.Errorf("Expected APP_NAME to be 'test-app', got '%s'", cfg.App.Name)
	}

	if cfg.HTTP.Port != "8080" {
		t.Errorf("Expected HTTP_PORT to be '8080', got '%s'", cfg.HTTP.Port)
	}

	if cfg.PG.PoolMax != 10 {
		t.Errorf("Expected PG_POOL_MAX to be 10, got %d", cfg.PG.PoolMax)
	}

	if !cfg.Metrics.Enabled {
		t.Error("Expected Metrics.Enabled to be true by default")
	}

	if cfg.Swagger.Enabled {
		t.Error("Expected Swagger.Enabled to be false by default")
	}

	os.Setenv("METRICS_ENABLED", "false")
	os.Setenv("SWAGGER_ENABLED", "true")

	cfg2, err := NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() failed: %v", err)
	}

	if cfg2.Metrics.Enabled {
		t.Error("Expected Metrics.Enabled to be false when set to 'false'")
	}

	if !cfg2.Swagger.Enabled {
		t.Error("Expected Swagger.Enabled to be true when set to 'true'")
	}

	os.Unsetenv("APP_NAME")
	os.Unsetenv("APP_VERSION")
	os.Unsetenv("HTTP_PORT")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("PG_URL")
	os.Unsetenv("PG_POOL_MAX")
	os.Unsetenv("RMQ_RPC_SERVER")
	os.Unsetenv("RMQ_RPC_CLIENT")
	os.Unsetenv("RMQ_URL")
	os.Unsetenv("METRICS_ENABLED")
	os.Unsetenv("SWAGGER_ENABLED")
}
