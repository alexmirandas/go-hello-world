package main

import (
	"testing"
)

// TestMainFunction is a simple test for the main function
func TestMainFunction(t *testing.T) {
	// Since main() prints output, we don't have a return value to assert.
	// Typically, we'd capture stdout if needed, but for this simple program, we just ensure it runs.
	main()
}
