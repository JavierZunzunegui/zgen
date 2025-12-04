package ziter_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestFilter tests the Filter function with iter.Seq
func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			filter:   func(int) bool { return true },
			expected: []int{},
		},
		{
			name:     "all match",
			input:    []int{1, 2, 3},
			filter:   func(int) bool { return true },
			expected: []int{1, 2, 3},
		},
		{
			name:     "none match",
			input:    []int{1, 2, 3},
			filter:   func(int) bool { return false },
			expected: []int{},
		},
		{
			name:     "even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			filter:   func(v int) bool { return v%2 == 0 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "odd numbers",
			input:    []int{1, 2, 3, 4, 5},
			filter:   func(v int) bool { return v%2 != 0 },
			expected: []int{1, 3, 5},
		},
		{
			name:     "single element match",
			input:    []int{5},
			filter:   func(v int) bool { return v == 5 },
			expected: []int{5},
		},
		{
			name:     "single element no match",
			input:    []int{5},
			filter:   func(v int) bool { return v == 10 },
			expected: []int{},
		},
		{
			name:     "greater than threshold",
			input:    []int{1, 5, 10, 15, 3, 8},
			filter:   func(v int) bool { return v > 7 },
			expected: []int{10, 15, 8},
		},
		{
			name:     "zero values included",
			input:    []int{0, 1, 2, 0, 3},
			filter:   func(v int) bool { return v == 0 },
			expected: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			filtered := ziter.Filter(seq, tt.filter)
			result := slices.Collect(filtered)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Filter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFilterEarlyTermination verifies Filter respects early termination
func TestFilterEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	filtered := ziter.Filter(input, func(v int) bool {
		callCount++
		return v%2 == 0
	})

	// Consume only 2 elements
	count := 0
	for range filtered {
		count++
		if count == 2 {
			break
		}
	}

	// The filter function should have been called for the first 4 elements
	// (to find 2 even numbers: 2 and 4)
	if callCount > 5 {
		t.Errorf("Filter didn't respect early termination: callCount=%d, expected <= 5", callCount)
	}
}

// TestFilter2 tests the Filter2 function with iter.Seq2
func TestFilter2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		filter   func(string, int) bool
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			filter:   func(string, int) bool { return true },
			expected: map[string]int{},
		},
		{
			name:     "all match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(string, int) bool { return true },
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "none match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(string, int) bool { return false },
			expected: map[string]int{},
		},
		{
			name:     "filter by value even",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter:   func(_ string, v int) bool { return v%2 == 0 },
			expected: map[string]int{"b": 2, "d": 4},
		},
		{
			name:     "filter by key prefix",
			input:    map[string]int{"apple": 1, "banana": 2, "apricot": 3, "cherry": 4},
			filter:   func(k string, _ int) bool { return len(k) > 0 && k[0] == 'a' },
			expected: map[string]int{"apple": 1, "apricot": 3},
		},
		{
			name:     "filter by both key and value",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			filter:   func(k string, v int) bool { return k >= "c" && v > 2 },
			expected: map[string]int{"c": 3, "d": 4, "e": 5},
		},
		{
			name:     "single entry match",
			input:    map[string]int{"x": 10},
			filter:   func(k string, v int) bool { return k == "x" && v == 10 },
			expected: map[string]int{"x": 10},
		},
		{
			name:     "single entry no match",
			input:    map[string]int{"x": 10},
			filter:   func(k string, v int) bool { return k == "y" },
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			filtered := ziter.Filter2(seq, tt.filter)
			result := maps.Collect(filtered)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("Filter2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFilter2EarlyTermination verifies Filter2 respects early termination
func TestFilter2EarlyTermination(t *testing.T) {
	callCount := 0
	input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	seq := maps.All(input)
	filtered := ziter.Filter2(seq, func(k string, v int) bool {
		callCount++
		return true
	})

	// Consume only 2 elements
	count := 0
	for range filtered {
		count++
		if count == 2 {
			break
		}
	}

	// The filter function should have been called approximately 2 times
	if callCount > 3 {
		t.Errorf("Filter2 didn't respect early termination: callCount=%d, expected <= 3", callCount)
	}
}

// TestFilterKey tests the FilterKey function
func TestFilterKey(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		filter   func(string) bool
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			filter:   func(string) bool { return true },
			expected: map[string]int{},
		},
		{
			name:     "all keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(string) bool { return true },
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "no keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(string) bool { return false },
			expected: map[string]int{},
		},
		{
			name:     "filter keys by prefix",
			input:    map[string]int{"apple": 1, "banana": 2, "apricot": 3, "cherry": 4},
			filter:   func(k string) bool { return len(k) > 0 && k[0] == 'a' },
			expected: map[string]int{"apple": 1, "apricot": 3},
		},
		{
			name:     "filter keys by length",
			input:    map[string]int{"a": 1, "ab": 2, "abc": 3, "abcd": 4},
			filter:   func(k string) bool { return len(k) > 2 },
			expected: map[string]int{"abc": 3, "abcd": 4},
		},
		{
			name:     "single key match",
			input:    map[string]int{"x": 10, "y": 20},
			filter:   func(k string) bool { return k == "x" },
			expected: map[string]int{"x": 10},
		},
		{
			name:     "preserve all values",
			input:    map[string]int{"keep": 100, "remove": 200, "keep2": 300},
			filter:   func(k string) bool { return k[0] == 'k' },
			expected: map[string]int{"keep": 100, "keep2": 300},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			filtered := ziter.FilterKey(seq, tt.filter)
			result := maps.Collect(filtered)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("FilterKey() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFilterValue tests the FilterValue function
func TestFilterValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		filter   func(int) bool
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			filter:   func(int) bool { return true },
			expected: map[string]int{},
		},
		{
			name:     "all values match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(int) bool { return true },
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "no values match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			filter:   func(int) bool { return false },
			expected: map[string]int{},
		},
		{
			name:     "filter even values",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			filter:   func(v int) bool { return v%2 == 0 },
			expected: map[string]int{"b": 2, "d": 4},
		},
		{
			name:     "filter values greater than threshold",
			input:    map[string]int{"small": 5, "medium": 10, "large": 15, "tiny": 2},
			filter:   func(v int) bool { return v > 7 },
			expected: map[string]int{"medium": 10, "large": 15},
		},
		{
			name:     "single value match",
			input:    map[string]int{"x": 10, "y": 20},
			filter:   func(v int) bool { return v == 10 },
			expected: map[string]int{"x": 10},
		},
		{
			name:     "preserve all keys",
			input:    map[string]int{"key1": 100, "key2": 50, "key3": 75},
			filter:   func(v int) bool { return v >= 75 },
			expected: map[string]int{"key1": 100, "key3": 75},
		},
		{
			name:     "zero values",
			input:    map[string]int{"a": 0, "b": 1, "c": 0, "d": 2},
			filter:   func(v int) bool { return v == 0 },
			expected: map[string]int{"a": 0, "c": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			filtered := ziter.FilterValue(seq, tt.filter)
			result := maps.Collect(filtered)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("FilterValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}
