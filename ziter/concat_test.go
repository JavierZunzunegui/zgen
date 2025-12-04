package ziter_test

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestConcat tests the Concat function
func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]int
		expected []int
	}{
		{
			name:     "no sequences",
			inputs:   [][]int{},
			expected: []int{},
		},
		{
			name:     "single empty sequence",
			inputs:   [][]int{{}},
			expected: []int{},
		},
		{
			name:     "single non-empty sequence",
			inputs:   [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "two sequences",
			inputs:   [][]int{{1, 2}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "three sequences",
			inputs:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty sequence at start",
			inputs:   [][]int{{}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty sequence in middle",
			inputs:   [][]int{{1, 2}, {}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "empty sequence at end",
			inputs:   [][]int{{1, 2, 3}, {}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all empty sequences",
			inputs:   [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "mixed empty and non-empty",
			inputs:   [][]int{{}, {1}, {}, {2, 3}, {}, {4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single element sequences",
			inputs:   [][]int{{1}, {2}, {3}, {4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "order preservation",
			inputs:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "varying sizes",
			inputs:   [][]int{{1}, {2, 3, 4}, {5, 6}, {7, 8, 9, 10}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert inputs to sequences
			seqs := make([]iter.Seq[int], len(tt.inputs))
			for i, input := range tt.inputs {
				seqs[i] = slices.Values(input)
			}

			// Concatenate
			concatenated := ziter.Concat(seqs...)
			result := slices.Collect(concatenated)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Concat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestConcat2 tests the Concat2 function
func TestConcat2(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []map[string]int
		validate func(map[string]int) bool
	}{
		{
			name:   "no sequences",
			inputs: []map[string]int{},
			validate: func(result map[string]int) bool {
				return len(result) == 0
			},
		},
		{
			name:   "single empty map",
			inputs: []map[string]int{{}},
			validate: func(result map[string]int) bool {
				return len(result) == 0
			},
		},
		{
			name:   "single non-empty map",
			inputs: []map[string]int{{"a": 1, "b": 2}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:   "two maps no overlap",
			inputs: []map[string]int{{"a": 1}, {"b": 2}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:   "two maps with overlap - last wins",
			inputs: []map[string]int{{"a": 1, "b": 2}, {"b": 20, "c": 3}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 20, "c": 3})
			},
		},
		{
			name:   "three maps",
			inputs: []map[string]int{{"a": 1}, {"b": 2}, {"c": 3}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2, "c": 3})
			},
		},
		{
			name:   "empty map at start",
			inputs: []map[string]int{{}, {"a": 1, "b": 2}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:   "empty map in middle",
			inputs: []map[string]int{{"a": 1}, {}, {"b": 2}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:   "empty map at end",
			inputs: []map[string]int{{"a": 1, "b": 2}, {}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2})
			},
		},
		{
			name:   "all empty maps",
			inputs: []map[string]int{{}, {}, {}},
			validate: func(result map[string]int) bool {
				return len(result) == 0
			},
		},
		{
			name:   "mixed empty and non-empty",
			inputs: []map[string]int{{}, {"a": 1}, {}, {"b": 2, "c": 3}, {}, {"d": 4}},
			validate: func(result map[string]int) bool {
				return maps.Equal(result, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert inputs to sequences
			seqs := make([]iter.Seq2[string, int], len(tt.inputs))
			for i, input := range tt.inputs {
				seqs[i] = maps.All(input)
			}

			// Concatenate
			concatenated := ziter.Concat2(seqs...)
			result := maps.Collect(concatenated)

			if !tt.validate(result) {
				t.Errorf("Concat2() = %v, validation failed", result)
			}
		})
	}
}

// TestConcatWithStrings tests Concat with string sequences
func TestConcatWithStrings(t *testing.T) {
	seq1 := slices.Values([]string{"hello", "world"})
	seq2 := slices.Values([]string{"foo", "bar"})
	seq3 := slices.Values([]string{"baz"})

	concatenated := ziter.Concat(seq1, seq2, seq3)
	result := slices.Collect(concatenated)
	expected := []string{"hello", "world", "foo", "bar", "baz"}

	if !slices.Equal(result, expected) {
		t.Errorf("Concat() = %v, want %v", result, expected)
	}
}
