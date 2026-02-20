package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

type kv struct {
	k, v int
}

func TestDedup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "consecutive duplicates",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "non-consecutive duplicates",
			input:    []int{1, 2, 1, 3, 2},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all duplicates",
			input:    []int{5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slices.Collect(ziter.Dedup(slices.Values(tt.input)))
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Dedup() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDedup2(t *testing.T) {
	// helper: build a Seq2[int,int] from a slice of kv pairs
	fromPairs := func(pairs []kv) func(func(int, int) bool) {
		return func(yield func(int, int) bool) {
			for _, p := range pairs {
				if !yield(p.k, p.v) {
					return
				}
			}
		}
	}
	collectPairs := func(seq func(func(int, int) bool)) []kv {
		var out []kv
		seq(func(k, v int) bool {
			out = append(out, kv{k, v})
			return true
		})
		return out
	}

	tests := []struct {
		name     string
		input    []kv
		expected []kv
	}{
		{
			name:     "empty",
			input:    []kv{},
			expected: nil,
		},
		{
			name:     "no duplicates",
			input:    []kv{{1, 1}, {1, 2}, {2, 1}},
			expected: []kv{{1, 1}, {1, 2}, {2, 1}},
		},
		{
			name:     "duplicate pair",
			input:    []kv{{1, 2}, {1, 2}, {3, 4}},
			expected: []kv{{1, 2}, {3, 4}},
		},
		{
			name:     "same key different value not duplicate",
			input:    []kv{{1, 2}, {1, 3}},
			expected: []kv{{1, 2}, {1, 3}},
		},
		{
			name:     "same value different key not duplicate",
			input:    []kv{{1, 2}, {3, 2}},
			expected: []kv{{1, 2}, {3, 2}},
		},
		{
			name:     "non-consecutive duplicate pair",
			input:    []kv{{1, 2}, {3, 4}, {1, 2}},
			expected: []kv{{1, 2}, {3, 4}},
		},
		{
			name:     "all same",
			input:    []kv{{5, 6}, {5, 6}, {5, 6}},
			expected: []kv{{5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := collectPairs(ziter.Dedup2(fromPairs(tt.input)))
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Dedup2() = %v, want %v", result, tt.expected)
			}
		})
	}
}
