package ziter_test

import (
	"maps"
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

// TestCount tests the Count function
func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "two elements",
			input:    []int{10, 20},
			expected: 2,
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "ten elements",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			result := ziter.Count(seq)

			if result != tt.expected {
				t.Errorf("Count() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestCountFullIteration verifies Count iterates through entire sequence
func TestCountFullIteration(t *testing.T) {
	callCount := 0
	input := func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			callCount++
			if !yield(i) {
				return
			}
		}
	}

	result := ziter.Count(input)

	if result != 10 {
		t.Errorf("Count() = %v, want 10", result)
	}

	if callCount != 10 {
		t.Errorf("Count should iterate through all elements, called %d times", callCount)
	}
}

// TestCount2 tests the Count2 function
func TestCount2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: 0,
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 1},
			expected: 1,
		},
		{
			name:     "two entries",
			input:    map[string]int{"a": 1, "b": 2},
			expected: 2,
		},
		{
			name:     "multiple entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			expected: 5,
		},
		{
			name:     "ten entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			result := ziter.Count2(seq)

			if result != tt.expected {
				t.Errorf("Count2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestCount2FullIteration verifies Count2 iterates through entire sequence
func TestCount2FullIteration(t *testing.T) {
	callCount := 0
	input := func(yield func(string, int) bool) {
		pairs := []struct {
			k string
			v int
		}{
			{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}, {"e", 5},
		}
		for _, p := range pairs {
			callCount++
			if !yield(p.k, p.v) {
				return
			}
		}
	}

	result := ziter.Count2(input)

	if result != 5 {
		t.Errorf("Count2() = %v, want 5", result)
	}

	if callCount != 5 {
		t.Errorf("Count2 should iterate through all elements, called %d times", callCount)
	}
}
