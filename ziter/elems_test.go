package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestExists tests the Exists function
func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			expected: false,
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "two elements",
			input:    []int{10, 20},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			result := ziter.Exists(seq)

			if result != tt.expected {
				t.Errorf("Exists() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestExistsEarlyTermination verifies Exists stops after checking first element
func TestExistsEarlyTermination(t *testing.T) {
	callCount := 0
	input := func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			callCount++
			if !yield(i) {
				return
			}
		}
	}

	result := ziter.Exists(input)

	if !result {
		t.Error("Exists() should return true for non-empty sequence")
	}

	if callCount != 1 {
		t.Errorf("Exists should check only first element, called %d times", callCount)
	}
}
