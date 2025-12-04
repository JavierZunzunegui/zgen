package ziter_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestSplit tests the Split function
func TestSplit(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		predicate     func(int) bool
		expectedTrue  []int
		expectedFalse []int
	}{
		{
			name:          "empty sequence",
			input:         []int{},
			predicate:     func(int) bool { return true },
			expectedTrue:  []int{},
			expectedFalse: []int{},
		},
		{
			name:          "all match",
			input:         []int{1, 2, 3},
			predicate:     func(int) bool { return true },
			expectedTrue:  []int{1, 2, 3},
			expectedFalse: []int{},
		},
		{
			name:          "none match",
			input:         []int{1, 2, 3},
			predicate:     func(int) bool { return false },
			expectedTrue:  []int{},
			expectedFalse: []int{1, 2, 3},
		},
		{
			name:          "even/odd split",
			input:         []int{1, 2, 3, 4, 5, 6},
			predicate:     func(v int) bool { return v%2 == 0 },
			expectedTrue:  []int{2, 4, 6},
			expectedFalse: []int{1, 3, 5},
		},
		{
			name:          "single element true",
			input:         []int{5},
			predicate:     func(v int) bool { return v == 5 },
			expectedTrue:  []int{5},
			expectedFalse: []int{},
		},
		{
			name:          "single element false",
			input:         []int{5},
			predicate:     func(v int) bool { return v == 10 },
			expectedTrue:  []int{},
			expectedFalse: []int{5},
		},
		{
			name:          "greater than threshold",
			input:         []int{1, 5, 10, 15, 3, 8},
			predicate:     func(v int) bool { return v > 7 },
			expectedTrue:  []int{10, 15, 8},
			expectedFalse: []int{1, 5, 3},
		},
		{
			name:          "positive/negative split",
			input:         []int{-2, 3, -1, 5, 0, -4, 7},
			predicate:     func(v int) bool { return v > 0 },
			expectedTrue:  []int{3, 5, 7},
			expectedFalse: []int{-2, -1, 0, -4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			trueSeq, falseSeq := ziter.Split(seq, tt.predicate)

			// Collect both sequences
			trueResult := slices.Collect(trueSeq)
			falseResult := slices.Collect(falseSeq)

			if !slices.Equal(trueResult, tt.expectedTrue) {
				t.Errorf("Split() true branch = %v, want %v", trueResult, tt.expectedTrue)
			}

			if !slices.Equal(falseResult, tt.expectedFalse) {
				t.Errorf("Split() false branch = %v, want %v", falseResult, tt.expectedFalse)
			}
		})
	}
}

// TestSplitIndependence verifies both sequences can be consumed independently
func TestSplitIndependence(t *testing.T) {
	input := slices.Values([]int{1, 2, 3, 4, 5, 6})
	trueSeq, falseSeq := ziter.Split(input, func(v int) bool { return v%2 == 0 })

	// Consume false sequence first
	falseResult := slices.Collect(falseSeq)
	if !slices.Equal(falseResult, []int{1, 3, 5}) {
		t.Errorf("false sequence = %v, want [1 3 5]", falseResult)
	}

	// Then consume true sequence
	trueResult := slices.Collect(trueSeq)
	if !slices.Equal(trueResult, []int{2, 4, 6}) {
		t.Errorf("true sequence = %v, want [2 4 6]", trueResult)
	}
}

// TestSplit2 tests the Split2 function
func TestSplit2(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(string, int) bool
		expectedTrue  map[string]int
		expectedFalse map[string]int
	}{
		{
			name:          "empty map",
			input:         map[string]int{},
			predicate:     func(string, int) bool { return true },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{},
		},
		{
			name:          "all match",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:     func(string, int) bool { return true },
			expectedTrue:  map[string]int{"a": 1, "b": 2, "c": 3},
			expectedFalse: map[string]int{},
		},
		{
			name:          "none match",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:     func(string, int) bool { return false },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:          "split by value",
			input:         map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate:     func(_ string, v int) bool { return v%2 == 0 },
			expectedTrue:  map[string]int{"b": 2, "d": 4},
			expectedFalse: map[string]int{"a": 1, "c": 3},
		},
		{
			name:          "split by key",
			input:         map[string]int{"apple": 1, "banana": 2, "apricot": 3, "cherry": 4},
			predicate:     func(k string, _ int) bool { return len(k) > 0 && k[0] == 'a' },
			expectedTrue:  map[string]int{"apple": 1, "apricot": 3},
			expectedFalse: map[string]int{"banana": 2, "cherry": 4},
		},
		{
			name:          "split by both key and value",
			input:         map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate:     func(k string, v int) bool { return k >= "c" && v > 2 },
			expectedTrue:  map[string]int{"c": 3, "d": 4, "e": 5},
			expectedFalse: map[string]int{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			trueSeq, falseSeq := ziter.Split2(seq, tt.predicate)

			// Collect both sequences
			trueResult := maps.Collect(trueSeq)
			falseResult := maps.Collect(falseSeq)

			if !maps.Equal(trueResult, tt.expectedTrue) {
				t.Errorf("Split2() true branch = %v, want %v", trueResult, tt.expectedTrue)
			}

			if !maps.Equal(falseResult, tt.expectedFalse) {
				t.Errorf("Split2() false branch = %v, want %v", falseResult, tt.expectedFalse)
			}
		})
	}
}

// TestSplitKey tests the SplitKey function
func TestSplitKey(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(string) bool
		expectedTrue  map[string]int
		expectedFalse map[string]int
	}{
		{
			name:          "empty map",
			input:         map[string]int{},
			predicate:     func(string) bool { return true },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{},
		},
		{
			name:          "all keys match",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(string) bool { return true },
			expectedTrue:  map[string]int{"a": 1, "b": 2},
			expectedFalse: map[string]int{},
		},
		{
			name:          "no keys match",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(string) bool { return false },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{"a": 1, "b": 2},
		},
		{
			name:          "split by key prefix",
			input:         map[string]int{"apple": 1, "banana": 2, "apricot": 3, "cherry": 4},
			predicate:     func(k string) bool { return len(k) > 0 && k[0] == 'a' },
			expectedTrue:  map[string]int{"apple": 1, "apricot": 3},
			expectedFalse: map[string]int{"banana": 2, "cherry": 4},
		},
		{
			name:          "split by key length",
			input:         map[string]int{"a": 1, "ab": 2, "abc": 3, "abcd": 4},
			predicate:     func(k string) bool { return len(k) > 2 },
			expectedTrue:  map[string]int{"abc": 3, "abcd": 4},
			expectedFalse: map[string]int{"a": 1, "ab": 2},
		},
		{
			name:          "values preserved in both",
			input:         map[string]int{"keep": 100, "remove": 200, "keep2": 300, "remove2": 400},
			predicate:     func(k string) bool { return k[0] == 'k' },
			expectedTrue:  map[string]int{"keep": 100, "keep2": 300},
			expectedFalse: map[string]int{"remove": 200, "remove2": 400},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			trueSeq, falseSeq := ziter.SplitKey(seq, tt.predicate)

			// Collect both sequences
			trueResult := maps.Collect(trueSeq)
			falseResult := maps.Collect(falseSeq)

			if !maps.Equal(trueResult, tt.expectedTrue) {
				t.Errorf("SplitKey() true branch = %v, want %v", trueResult, tt.expectedTrue)
			}

			if !maps.Equal(falseResult, tt.expectedFalse) {
				t.Errorf("SplitKey() false branch = %v, want %v", falseResult, tt.expectedFalse)
			}
		})
	}
}

// TestSplitValue tests the SplitValue function
func TestSplitValue(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(int) bool
		expectedTrue  map[string]int
		expectedFalse map[string]int
	}{
		{
			name:          "empty map",
			input:         map[string]int{},
			predicate:     func(int) bool { return true },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{},
		},
		{
			name:          "all values match",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(int) bool { return true },
			expectedTrue:  map[string]int{"a": 1, "b": 2},
			expectedFalse: map[string]int{},
		},
		{
			name:          "no values match",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(int) bool { return false },
			expectedTrue:  map[string]int{},
			expectedFalse: map[string]int{"a": 1, "b": 2},
		},
		{
			name:          "split even/odd values",
			input:         map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate:     func(v int) bool { return v%2 == 0 },
			expectedTrue:  map[string]int{"b": 2, "d": 4},
			expectedFalse: map[string]int{"a": 1, "c": 3, "e": 5},
		},
		{
			name:          "split by value threshold",
			input:         map[string]int{"small": 5, "medium": 10, "large": 15, "tiny": 2},
			predicate:     func(v int) bool { return v > 7 },
			expectedTrue:  map[string]int{"medium": 10, "large": 15},
			expectedFalse: map[string]int{"small": 5, "tiny": 2},
		},
		{
			name:          "keys preserved in both",
			input:         map[string]int{"key1": 100, "key2": 50, "key3": 75, "key4": 200},
			predicate:     func(v int) bool { return v >= 75 },
			expectedTrue:  map[string]int{"key1": 100, "key3": 75, "key4": 200},
			expectedFalse: map[string]int{"key2": 50},
		},
		{
			name:          "zero values",
			input:         map[string]int{"a": 0, "b": 1, "c": 0, "d": 2},
			predicate:     func(v int) bool { return v == 0 },
			expectedTrue:  map[string]int{"a": 0, "c": 0},
			expectedFalse: map[string]int{"b": 1, "d": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			trueSeq, falseSeq := ziter.SplitValue(seq, tt.predicate)

			// Collect both sequences
			trueResult := maps.Collect(trueSeq)
			falseResult := maps.Collect(falseSeq)

			if !maps.Equal(trueResult, tt.expectedTrue) {
				t.Errorf("SplitValue() true branch = %v, want %v", trueResult, tt.expectedTrue)
			}

			if !maps.Equal(falseResult, tt.expectedFalse) {
				t.Errorf("SplitValue() false branch = %v, want %v", falseResult, tt.expectedFalse)
			}
		})
	}
}

// TestSplitEarlyTermination verifies early termination on one sequence doesn't affect the other
func TestSplitEarlyTermination(t *testing.T) {
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	trueSeq, falseSeq := ziter.Split(input, func(v int) bool { return v%2 == 0 })

	// Consume only 2 elements from true sequence
	count := 0
	for range trueSeq {
		count++
		if count == 2 {
			break
		}
	}

	// Should still be able to consume entire false sequence
	falseResult := slices.Collect(falseSeq)
	expectedFalse := []int{1, 3, 5, 7, 9}

	if !slices.Equal(falseResult, expectedFalse) {
		t.Errorf("false sequence after early termination of true = %v, want %v", falseResult, expectedFalse)
	}
}
