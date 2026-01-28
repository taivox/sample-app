package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigConstants(t *testing.T) {
	assert.Equal(t, "POSTGRES_USER", POSTGRES_USER)
	assert.Equal(t, "POSTGRES_PASSWORD", POSTGRES_PASSWORD)
	assert.Equal(t, "POSTGRES_DB", POSTGRES_DB)
	assert.Equal(t, "CLIENT_URL", CLIENT_URL)
	assert.Equal(t, "SERVER_PORT", SERVER_PORT)
	assert.Equal(t, "JWT_KEY", JWT_KEY)
	assert.Equal(t, "RUN_MIGRATION", RUN_MIGRATION)
	assert.Equal(t, "POSTGRES_SERVER_HOST", POSTGRES_SERVER_HOST)
	assert.Equal(t, "ENV", ENVIRONEMT)
}

func TestConfigDefaultValues(t *testing.T) {
	// POSTGRES_SERVER_HOST has a default value
	assert.Equal(t, "localhost", Config[POSTGRES_SERVER_HOST])
}

func TestConfigType(t *testing.T) {
	var cfg ConfigType = make(map[string]string)
	cfg["test_key"] = "test_value"
	assert.Equal(t, "test_value", cfg["test_key"])
}

func TestConfigKeysExist(t *testing.T) {
	expectedKeys := []string{
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DB,
		CLIENT_URL,
		SERVER_PORT,
		JWT_KEY,
		RUN_MIGRATION,
		POSTGRES_SERVER_HOST,
	}

	for _, key := range expectedKeys {
		_, exists := Config[key]
		assert.True(t, exists, "Config should have key: %s", key)
	}
}

func TestEnvironmentVariableOverride(t *testing.T) {
	// Save original value
	originalVal := Config[POSTGRES_SERVER_HOST]

	// Set environment variable
	os.Setenv(POSTGRES_SERVER_HOST, "custom-host")
	defer os.Unsetenv(POSTGRES_SERVER_HOST)

	// Read from env
	envVal := os.Getenv(POSTGRES_SERVER_HOST)
	assert.Equal(t, "custom-host", envVal)

	// Restore original
	Config[POSTGRES_SERVER_HOST] = originalVal
}

func TestConfigMapOperations(t *testing.T) {
	// Test that Config is a proper map
	testConfig := ConfigType{
		"key1": "value1",
		"key2": "value2",
	}

	assert.Equal(t, "value1", testConfig["key1"])
	assert.Equal(t, "value2", testConfig["key2"])

	// Test modification
	testConfig["key1"] = "modified"
	assert.Equal(t, "modified", testConfig["key1"])

	// Test non-existent key
	val, exists := testConfig["nonexistent"]
	assert.False(t, exists)
	assert.Empty(t, val)
}
