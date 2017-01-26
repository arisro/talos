package config

import "testing"

func TestHello(t *testing.T) {
	if val := Hello(); val != "hello123" {
		t.Fatalf("Expected 'hello', got %s instead.", val)
	}
}
