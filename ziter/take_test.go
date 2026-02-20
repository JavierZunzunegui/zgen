package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

func TestTake(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "take from empty",
			input:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "take zero",
			input:    []int{1, 2, 3},
			n:        0,
			expected: []int{},
		},
		{
			name:     "take negative",
			input:    []int{1, 2, 3},
			n:        -1,
			expected: []int{},
		},
		{
			name:     "take fewer than available",
			input:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take exact count",
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take more than available",
			input:    []int{1, 2},
			n:        5,
			expected: []int{1, 2},
		},
		{
			name:     "take one",
			input:    []int{10, 20, 30},
			n:        1,
			expected: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slices.Collect(ziter.Take(slices.Values(tt.input), tt.n))
			if result == nil {
				result = []int{}
			}
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Take(%v, %d) = %v, want %v", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTakeEarlyTermination(t *testing.T) {
	callCount := 0
	input := func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			callCount++
			if !yield(i) {
				return
			}
		}
	}

	result := slices.Collect(ziter.Take(input, 3))
	expected := []int{1, 2, 3}

	if !slices.Equal(result, expected) {
		t.Errorf("Take() = %v, want %v", result, expected)
	}
	if callCount != 3 {
		t.Errorf("Take should only consume 3 elements, consumed %d", callCount)
	}
}

func TestTake2(t *testing.T) {
	tests := []struct {
		name       string
		keys       []string
		values     []int
		n          int
		wantKeys   []string
		wantValues []int
	}{
		{
			name:       "take from empty",
			keys:       []string{},
			values:     []int{},
			n:          3,
			wantKeys:   []string{},
			wantValues: []int{},
		},
		{
			name:       "take zero",
			keys:       []string{"a", "b", "c"},
			values:     []int{1, 2, 3},
			n:          0,
			wantKeys:   []string{},
			wantValues: []int{},
		},
		{
			name:       "take fewer than available",
			keys:       []string{"a", "b", "c", "d"},
			values:     []int{1, 2, 3, 4},
			n:          2,
			wantKeys:   []string{"a", "b"},
			wantValues: []int{1, 2},
		},
		{
			name:       "take more than available",
			keys:       []string{"a", "b"},
			values:     []int{1, 2},
			n:          5,
			wantKeys:   []string{"a", "b"},
			wantValues: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := makeSeq2(tt.keys, tt.values)
			gotKeys, gotValues := collectSeq2(ziter.Take2(seq, tt.n))
			if gotKeys == nil {
				gotKeys = []string{}
			}
			if gotValues == nil {
				gotValues = []int{}
			}
			if !slices.Equal(gotKeys, tt.wantKeys) || !slices.Equal(gotValues, tt.wantValues) {
				t.Errorf("Take2(%v, %v, %d) = (%v, %v), want (%v, %v)",
					tt.keys, tt.values, tt.n, gotKeys, gotValues, tt.wantKeys, tt.wantValues)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "drop from empty",
			input:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop zero",
			input:    []int{1, 2, 3},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop negative",
			input:    []int{1, 2, 3},
			n:        -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop fewer than available",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5},
		},
		{
			name:     "drop exact count",
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop more than available",
			input:    []int{1, 2},
			n:        5,
			expected: []int{},
		},
		{
			name:     "drop one",
			input:    []int{10, 20, 30},
			n:        1,
			expected: []int{20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slices.Collect(ziter.Drop(slices.Values(tt.input), tt.n))
			if result == nil {
				result = []int{}
			}
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Drop(%v, %d) = %v, want %v", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDrop2(t *testing.T) {
	tests := []struct {
		name       string
		keys       []string
		values     []int
		n          int
		wantKeys   []string
		wantValues []int
	}{
		{
			name:       "drop from empty",
			keys:       []string{},
			values:     []int{},
			n:          3,
			wantKeys:   []string{},
			wantValues: []int{},
		},
		{
			name:       "drop zero",
			keys:       []string{"a", "b", "c"},
			values:     []int{1, 2, 3},
			n:          0,
			wantKeys:   []string{"a", "b", "c"},
			wantValues: []int{1, 2, 3},
		},
		{
			name:       "drop fewer than available",
			keys:       []string{"a", "b", "c", "d"},
			values:     []int{1, 2, 3, 4},
			n:          2,
			wantKeys:   []string{"c", "d"},
			wantValues: []int{3, 4},
		},
		{
			name:       "drop more than available",
			keys:       []string{"a", "b"},
			values:     []int{1, 2},
			n:          5,
			wantKeys:   []string{},
			wantValues: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := makeSeq2(tt.keys, tt.values)
			gotKeys, gotValues := collectSeq2(ziter.Drop2(seq, tt.n))
			if gotKeys == nil {
				gotKeys = []string{}
			}
			if gotValues == nil {
				gotValues = []int{}
			}
			if !slices.Equal(gotKeys, tt.wantKeys) || !slices.Equal(gotValues, tt.wantValues) {
				t.Errorf("Drop2(%v, %v, %d) = (%v, %v), want (%v, %v)",
					tt.keys, tt.values, tt.n, gotKeys, gotValues, tt.wantKeys, tt.wantValues)
			}
		})
	}
}

// helpers for Seq2 tests

func makeSeq2[K, V any](keys []K, values []V) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for i := range keys {
			if !yield(keys[i], values[i]) {
				return
			}
		}
	}
}

func collectSeq2[K, V any](seq func(func(K, V) bool)) ([]K, []V) {
	var ks []K
	var vs []V
	seq(func(k K, v V) bool {
		ks = append(ks, k)
		vs = append(vs, v)
		return true
	})
	return ks, vs
}
