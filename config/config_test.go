package config

import "testing"

func TestHello(t *testing.T) {
	if val := Hello(); val != "hello" {
		t.Fatalf("Expected 'hello', got %s instead.", val)
	}
}
