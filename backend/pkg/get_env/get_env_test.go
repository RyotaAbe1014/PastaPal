package get_env

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	// Test case when environment variable is set
	os.Setenv("TEST_KEY", "test_value")
	value := Getenv("TEST_KEY", "default_value")
	if value != "test_value" {
		t.Errorf("expected 'test_value', got '%s'", value)
	}
	os.Unsetenv("TEST_KEY")

	// Test case when environment variable is not set
	value = Getenv("NON_EXISTENT_KEY", "default_value")
	if value != "default_value" {
		t.Errorf("expected 'default_value', got '%s'", value)
	}
}
