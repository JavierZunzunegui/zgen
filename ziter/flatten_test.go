package ziter_test

import (
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestFlatten tests the Flatten function
func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		flatFunc func(string) []rune
		expected []rune
	}{
		{
			name:     "empty sequence",
			input:    []string{},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{},
		},
		{
			name:     "single element to empty",
			input:    []string{""},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{},
		},
		{
			name:     "single element to single",
			input:    []string{"a"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a'},
		},
		{
			name:     "single element to multiple",
			input:    []string{"abc"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a', 'b', 'c'},
		},
		{
			name:     "multiple elements to single each",
			input:    []string{"a", "b", "c"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a', 'b', 'c'},
		},
		{
			name:     "multiple elements to multiple each",
			input:    []string{"ab", "cd", "ef"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a', 'b', 'c', 'd', 'e', 'f'},
		},
		{
			name:     "mixed sizes",
			input:    []string{"a", "bc", "", "def"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a', 'b', 'c', 'd', 'e', 'f'},
		},
		{
			name:     "all flatten to empty",
			input:    []string{"a", "b", "c"},
			flatFunc: func(string) []rune { return []rune{} },
			expected: []rune{},
		},
		{
			name:     "some flatten to empty",
			input:    []string{"a", "", "b", "", "c"},
			flatFunc: func(s string) []rune { return []rune(s) },
			expected: []rune{'a', 'b', 'c'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			flattened := ziter.Flatten(seq, tt.flatFunc)
			result := slices.Collect(flattened)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Flatten() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFlattenInts tests Flatten with integer expansion
func TestFlattenInts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		flatFunc func(int) []int
		expected []int
	}{
		{
			name:  "duplicate each element",
			input: []int{1, 2, 3},
			flatFunc: func(v int) []int {
				return []int{v, v}
			},
			expected: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:  "expand to range",
			input: []int{1, 2, 3},
			flatFunc: func(v int) []int {
				result := make([]int, v)
				for i := range result {
					result[i] = v
				}
				return result
			},
			expected: []int{1, 2, 2, 3, 3, 3},
		},
		{
			name:  "conditional expansion",
			input: []int{1, 2, 3, 4, 5},
			flatFunc: func(v int) []int {
				if v%2 == 0 {
					return []int{v, v * 10}
				}
				return []int{v}
			},
			expected: []int{1, 2, 20, 3, 4, 40, 5},
		},
		{
			name:  "empty result from some",
			input: []int{1, 2, 3, 4},
			flatFunc: func(v int) []int {
				if v%2 == 0 {
					return []int{}
				}
				return []int{v}
			},
			expected: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			flattened := ziter.Flatten(seq, tt.flatFunc)
			result := slices.Collect(flattened)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Flatten() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFlattenEarlyTermination verifies Flatten respects early termination
func TestFlattenEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]string{"abc", "def", "ghi"})

	flattened := ziter.Flatten(input, func(s string) []rune {
		callCount++
		return []rune(s)
	})

	// Consume only 4 elements (should process "abc" and part of "def")
	count := 0
	for range flattened {
		count++
		if count == 4 {
			break
		}
	}

	// Should have called flatFunc for "abc" and "def"
	if callCount > 3 {
		t.Errorf("Flatten didn't respect early termination: callCount=%d, expected <= 3", callCount)
	}
}

// TestFlattenKeys tests the FlattenKeys function
func TestFlattenKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		flatFunc func(string) []string
		validate func(map[string]int) bool
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			flatFunc: func(s string) []string { return []string{s, s + s} },
			validate: func(result map[string]int) bool {
				return len(result) == 0
			},
		},
		{
			name:  "single key to single",
			input: map[string]int{"a": 1},
			flatFunc: func(s string) []string {
				return []string{s}
			},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1})
			},
		},
		{
			name:  "single key to multiple",
			input: map[string]int{"a": 1},
			flatFunc: func(s string) []string {
				return []string{s + "1", s + "2", s + "3"}
			},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a1": 1, "a2": 1, "a3": 1})
			},
		},
		{
			name:  "multiple keys to single each",
			input: map[string]int{"a": 1, "b": 2},
			flatFunc: func(s string) []string {
				return []string{s}
			},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:  "multiple keys to multiple each",
			input: map[string]int{"a": 1, "b": 2},
			flatFunc: func(s string) []string {
				return []string{s + "1", s + "2"}
			},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{
					"a1": 1, "a2": 1, "b1": 2, "b2": 2,
				})
			},
		},
		{
			name:  "some keys to empty",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			flatFunc: func(s string) []string {
				if s == "b" {
					return []string{}
				}
				return []string{s}
			},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "c": 3})
			},
		},
		{
			name:  "value duplication",
			input: map[string]int{"x": 10},
			flatFunc: func(s string) []string {
				return []string{s + "1", s + "2", s + "3"}
			},
			validate: func(result map[string]int) bool {
				// Same value (10) should be paired with all three keys
				return maps.Equal(result, map[string]int{"x1": 10, "x2": 10, "x3": 10})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			flattened := ziter.FlattenKeys(seq, tt.flatFunc)
			result := maps.Collect(flattened)

			if !tt.validate(result) {
				t.Errorf("FlattenKeys() = %v, validation failed", result)
			}
		})
	}
}

// TestFlattenKeysWithInts tests FlattenKeys with integer keys
func TestFlattenKeysWithInts(t *testing.T) {
	input := maps.All(map[int]string{1: "a", 2: "b"})

	flattened := ziter.FlattenKeys(input, func(k int) []int {
		return []int{k * 10, k * 100}
	})

	result := maps.Collect(flattened)
	expected := map[int]string{
		10: "a", 100: "a",
		20: "b", 200: "b",
	}

	if !maps.Equal(result, expected) {
		t.Errorf("FlattenKeys() = %v, want %v", result, expected)
	}
}

// TestFlattenValues tests the FlattenValues function
func TestFlattenValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		flatFunc func(string) []rune
		validate func(map[string]rune) bool
	}{
		{
			name:     "empty map",
			input:    map[string]string{},
			flatFunc: func(s string) []rune { return []rune(s) },
			validate: func(result map[string]rune) bool {
				return len(result) == 0
			},
		},
		{
			name:  "single value to single",
			input: map[string]string{"key": "a"},
			flatFunc: func(s string) []rune {
				return []rune(s)
			},
			validate: func(result map[string]rune) bool {
				return maps.Equal(result, map[string]rune{"key": 'a'})
			},
		},
		{
			name:  "single value to multiple",
			input: map[string]string{"key": "abc"},
			flatFunc: func(s string) []rune {
				return []rune(s)
			},
			validate: func(result map[string]rune) bool {
				// Same key paired with multiple values - last wins
				return result["key"] == 'c' && len(result) == 1
			},
		},
		{
			name:  "multiple values to single each",
			input: map[string]string{"k1": "a", "k2": "b"},
			flatFunc: func(s string) []rune {
				return []rune(s)
			},
			validate: func(result map[string]rune) bool {
				return maps.Equal(result, map[string]rune{"k1": 'a', "k2": 'b'})
			},
		},
		{
			name:  "some values to empty",
			input: map[string]string{"k1": "a", "k2": "", "k3": "b"},
			flatFunc: func(s string) []rune {
				return []rune(s)
			},
			validate: func(result map[string]rune) bool {
				// k2 should be missing because its value flattens to empty
				return maps.Equal(result, map[string]rune{"k1": 'a', "k3": 'b'})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			flattened := ziter.FlattenValues(seq, tt.flatFunc)
			result := maps.Collect(flattened)

			if !tt.validate(result) {
				t.Errorf("FlattenValues() = %v, validation failed", result)
			}
		})
	}
}

// TestFlattenValuesWithInts tests FlattenValues with integer expansion
func TestFlattenValuesWithInts(t *testing.T) {
	input := maps.All(map[string]int{"a": 1, "b": 2})

	flattened := ziter.FlattenValues(input, func(v int) []int {
		return []int{v, v * 10}
	})

	result := maps.Collect(flattened)

	// Each key will have multiple entries, last wins
	// "a" will have pairs: ("a", 1), ("a", 10) -> final: ("a", 10)
	// "b" will have pairs: ("b", 2), ("b", 20) -> final: ("b", 20)
	expected := map[string]int{"a": 10, "b": 20}

	if !maps.Equal(result, expected) {
		t.Errorf("FlattenValues() = %v, want %v", result, expected)
	}
}

// TestFlattenValuesKeyDuplication tests that keys are properly duplicated
func TestFlattenValuesKeyDuplication(t *testing.T) {
	input := maps.All(map[string]int{"x": 3})

	// Track all emitted pairs
	var pairs []struct {
		k string
		v int
	}

	flattened := ziter.FlattenValues(input, func(v int) []int {
		result := make([]int, v)
		for i := range result {
			result[i] = i + 1
		}
		return result
	})

	// Collect all pairs in order
	for k, v := range flattened {
		pairs = append(pairs, struct {
			k string
			v int
		}{k, v})
	}

	// Should have 3 pairs all with key "x"
	if len(pairs) != 3 {
		t.Errorf("expected 3 pairs, got %d", len(pairs))
	}

	for _, p := range pairs {
		if p.k != "x" {
			t.Errorf("expected key 'x', got '%s'", p.k)
		}
	}
}

// TestFlattenWordSplit tests Flatten with word splitting
func TestFlattenWordSplit(t *testing.T) {
	input := slices.Values([]string{"hello world", "foo bar baz"})

	flattened := ziter.Flatten(input, func(s string) []string {
		return strings.Fields(s)
	})

	result := slices.Collect(flattened)
	expected := []string{"hello", "world", "foo", "bar", "baz"}

	if !slices.Equal(result, expected) {
		t.Errorf("Flatten() = %v, want %v", result, expected)
	}
}
