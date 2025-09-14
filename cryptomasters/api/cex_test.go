package api_test

import (
	"Go/cryptomasters/api"
	"testing"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected error for empty currency code, got nil")
	}
}
