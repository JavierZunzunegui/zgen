package ziter_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		f       func(int, int) int
		wantVal int
		wantOk  bool
	}{
		{
			name:    "empty sequence",
			input:   []int{},
			f:       func(a, b int) int { return a + b },
			wantVal: 0,
			wantOk:  false,
		},
		{
			name:    "single element",
			input:   []int{42},
			f:       func(a, b int) int { return a + b },
			wantVal: 42,
			wantOk:  true,
		},
		{
			name:    "sum",
			input:   []int{1, 2, 3, 4, 5},
			f:       func(a, b int) int { return a + b },
			wantVal: 15,
			wantOk:  true,
		},
		{
			name:    "product",
			input:   []int{2, 3, 4},
			f:       func(a, b int) int { return a * b },
			wantVal: 24,
			wantOk:  true,
		},
		{
			name:  "max via reduce",
			input: []int{3, 7, 2, 9, 1},
			f: func(a, b int) int {
				if b > a {
					return b
				}
				return a
			},
			wantVal: 9,
			wantOk:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ziter.Reduce(slices.Values(tt.input), tt.f)
			if ok != tt.wantOk {
				t.Errorf("Reduce() ok = %v, want %v", ok, tt.wantOk)
			}
			if got != tt.wantVal {
				t.Errorf("Reduce() = %v, want %v", got, tt.wantVal)
			}
		})
	}
}

func TestReduce2(t *testing.T) {
	tests := []struct {
		name    string
		keys    []string
		values  []int
		f       func(string, int, string, int) (string, int)
		wantKey string
		wantVal int
		wantOk  bool
	}{
		{
			name:    "empty sequence",
			keys:    []string{},
			values:  []int{},
			f:       func(ak string, av int, bk string, bv int) (string, int) { return ak, av + bv },
			wantKey: "",
			wantVal: 0,
			wantOk:  false,
		},
		{
			name:    "single element",
			keys:    []string{"a"},
			values:  []int{1},
			f:       func(ak string, av int, bk string, bv int) (string, int) { return ak, av + bv },
			wantKey: "a",
			wantVal: 1,
			wantOk:  true,
		},
		{
			name:    "sum values keep first key",
			keys:    []string{"a", "b", "c"},
			values:  []int{10, 20, 30},
			f:       func(ak string, av int, bk string, bv int) (string, int) { return ak, av + bv },
			wantKey: "a",
			wantVal: 60,
			wantOk:  true,
		},
		{
			name:   "pick max value pair",
			keys:   []string{"a", "b", "c"},
			values: []int{10, 30, 20},
			f: func(ak string, av int, bk string, bv int) (string, int) {
				if bv > av {
					return bk, bv
				}
				return ak, av
			},
			wantKey: "b",
			wantVal: 30,
			wantOk:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := makeSeq2(tt.keys, tt.values)
			gotK, gotV, ok := ziter.Reduce2(seq, tt.f)
			if ok != tt.wantOk {
				t.Errorf("Reduce2() ok = %v, want %v", ok, tt.wantOk)
			}
			if gotK != tt.wantKey {
				t.Errorf("Reduce2() key = %v, want %v", gotK, tt.wantKey)
			}
			if gotV != tt.wantVal {
				t.Errorf("Reduce2() value = %v, want %v", gotV, tt.wantVal)
			}
		})
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

func intCmp(a, b int) int { return a - b }

func TestMaxFunc(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantVal int
		wantOk  bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{7}, 7, true},
		{"max at start", []int{9, 3, 5}, 9, true},
		{"max at middle", []int{1, 9, 2}, 9, true},
		{"max at end", []int{1, 3, 9}, 9, true},
		{"all equal", []int{4, 4, 4}, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ziter.MaxFunc(slices.Values(tt.input), intCmp)
			if ok != tt.wantOk || got != tt.wantVal {
				t.Errorf("Max(%v) = (%v, %v), want (%v, %v)", tt.input, got, ok, tt.wantVal, tt.wantOk)
			}
		})
	}
}

func TestMinFunc(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantVal int
		wantOk  bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{7}, 7, true},
		{"min at start", []int{1, 5, 9}, 1, true},
		{"min at middle", []int{5, 1, 9}, 1, true},
		{"min at end", []int{9, 5, 1}, 1, true},
		{"all equal", []int{4, 4, 4}, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ziter.MinFunc(slices.Values(tt.input), intCmp)
			if ok != tt.wantOk || got != tt.wantVal {
				t.Errorf("Min(%v) = (%v, %v), want (%v, %v)", tt.input, got, ok, tt.wantVal, tt.wantOk)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantVal int
		wantOk  bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{7}, 7, true},
		{"max at start", []int{9, 3, 5}, 9, true},
		{"max at middle", []int{1, 9, 2}, 9, true},
		{"max at end", []int{1, 3, 9}, 9, true},
		{"all equal", []int{4, 4, 4}, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ziter.Max(slices.Values(tt.input))
			if ok != tt.wantOk || got != tt.wantVal {
				t.Errorf("Max(%v) = (%v, %v), want (%v, %v)", tt.input, got, ok, tt.wantVal, tt.wantOk)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		wantVal int
		wantOk  bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{7}, 7, true},
		{"min at start", []int{1, 5, 9}, 1, true},
		{"min at middle", []int{5, 1, 9}, 1, true},
		{"min at end", []int{9, 5, 1}, 1, true},
		{"all equal", []int{4, 4, 4}, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ziter.Min(slices.Values(tt.input))
			if ok != tt.wantOk || got != tt.wantVal {
				t.Errorf("Min(%v) = (%v, %v), want (%v, %v)", tt.input, got, ok, tt.wantVal, tt.wantOk)
			}
		})
	}
}
