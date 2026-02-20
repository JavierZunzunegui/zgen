package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen"
	"github.com/JavierZunzunegui/zgen/ziter"
)

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
	// helper: build a Seq2[int,int] from a slice of Pair[int,int]
	fromPairs := func(pairs []zgen.Pair[int, int]) func(func(int, int) bool) {
		return func(yield func(int, int) bool) {
			for _, p := range pairs {
				if !yield(p.Both()) {
					return
				}
			}
		}
	}
	collectPairs := func(seq func(func(int, int) bool)) []zgen.Pair[int, int] {
		var out []zgen.Pair[int, int]
		seq(func(k, v int) bool {
			out = append(out, zgen.NewPair(k, v))
			return true
		})
		return out
	}

	tests := []struct {
		name     string
		input    []zgen.Pair[int, int]
		expected []zgen.Pair[int, int]
	}{
		{
			name:     "empty",
			input:    []zgen.Pair[int, int]{},
			expected: nil,
		},
		{
			name:     "no duplicates",
			input:    []zgen.Pair[int, int]{zgen.NewPair(1, 1), zgen.NewPair(1, 2), zgen.NewPair(2, 1)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(1, 1), zgen.NewPair(1, 2), zgen.NewPair(2, 1)},
		},
		{
			name:     "duplicate pair",
			input:    []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(1, 2), zgen.NewPair(3, 4)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(3, 4)},
		},
		{
			name:     "same key different value not duplicate",
			input:    []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(1, 3)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(1, 3)},
		},
		{
			name:     "same value different key not duplicate",
			input:    []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(3, 2)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(3, 2)},
		},
		{
			name:     "non-consecutive duplicate pair",
			input:    []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(3, 4), zgen.NewPair(1, 2)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(1, 2), zgen.NewPair(3, 4)},
		},
		{
			name:     "all same",
			input:    []zgen.Pair[int, int]{zgen.NewPair(5, 6), zgen.NewPair(5, 6), zgen.NewPair(5, 6)},
			expected: []zgen.Pair[int, int]{zgen.NewPair(5, 6)},
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
