package ziter_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestFindAny tests the FindAny function
func TestFindAny(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expectOk  bool
		validator func(int) bool // validates the returned value is acceptable
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			expectOk: false,
		},
		{
			name:      "single element",
			input:     []int{42},
			expectOk:  true,
			validator: func(v int) bool { return v == 42 },
		},
		{
			name:      "multiple elements",
			input:     []int{1, 2, 3, 4, 5},
			expectOk:  true,
			validator: func(v int) bool { return v >= 1 && v <= 5 },
		},
		{
			name:      "two elements",
			input:     []int{10, 20},
			expectOk:  true,
			validator: func(v int) bool { return v == 10 || v == 20 },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			result, ok := ziter.FindAny(seq)

			if ok != tt.expectOk {
				t.Errorf("FindAny() ok = %v, want %v", ok, tt.expectOk)
			}

			if ok && tt.validator != nil && !tt.validator(result) {
				t.Errorf("FindAny() returned unexpected value = %v", result)
			}

			if !ok && result != 0 {
				t.Errorf("FindAny() should return zero value when ok=false, got %v", result)
			}
		})
	}
}

// TestFindAnyEarlyTermination verifies FindAny stops after first element
func TestFindAnyEarlyTermination(t *testing.T) {
	callCount := 0
	input := func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			callCount++
			if !yield(i) {
				return
			}
		}
	}

	result, ok := ziter.FindAny(input)

	if !ok {
		t.Fatal("FindAny() should find an element")
	}

	if result != 1 {
		t.Errorf("FindAny() = %v, want 1", result)
	}

	if callCount != 1 {
		t.Errorf("FindAny should call yield exactly once, called %d times", callCount)
	}
}

// TestFindFirst tests the FindFirst function
func TestFindFirst(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		predicate  func(int) bool
		expected   int
		expectedOk bool
	}{
		{
			name:       "empty sequence",
			input:      []int{},
			predicate:  func(int) bool { return true },
			expected:   0,
			expectedOk: false,
		},
		{
			name:       "no match",
			input:      []int{1, 2, 3, 4, 5},
			predicate:  func(v int) bool { return v > 10 },
			expected:   0,
			expectedOk: false,
		},
		{
			name:       "first element matches",
			input:      []int{1, 2, 3, 4, 5},
			predicate:  func(v int) bool { return v == 1 },
			expected:   1,
			expectedOk: true,
		},
		{
			name:       "middle element matches",
			input:      []int{1, 2, 3, 4, 5},
			predicate:  func(v int) bool { return v == 3 },
			expected:   3,
			expectedOk: true,
		},
		{
			name:       "last element matches",
			input:      []int{1, 2, 3, 4, 5},
			predicate:  func(v int) bool { return v == 5 },
			expected:   5,
			expectedOk: true,
		},
		{
			name:       "multiple matches returns first",
			input:      []int{2, 4, 6, 8, 10},
			predicate:  func(v int) bool { return v%2 == 0 },
			expected:   2,
			expectedOk: true,
		},
		{
			name:       "single element match",
			input:      []int{42},
			predicate:  func(v int) bool { return v == 42 },
			expected:   42,
			expectedOk: true,
		},
		{
			name:       "single element no match",
			input:      []int{42},
			predicate:  func(v int) bool { return v == 10 },
			expected:   0,
			expectedOk: false,
		},
		{
			name:       "find greater than threshold",
			input:      []int{1, 3, 5, 7, 9, 11},
			predicate:  func(v int) bool { return v > 6 },
			expected:   7,
			expectedOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			result, ok := ziter.FindFirst(seq, tt.predicate)

			if ok != tt.expectedOk {
				t.Errorf("FindFirst() ok = %v, want %v", ok, tt.expectedOk)
			}

			if result != tt.expected {
				t.Errorf("FindFirst() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestFindFirstEarlyTermination verifies FindFirst stops after finding first match
func TestFindFirstEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result, ok := ziter.FindFirst(input, func(v int) bool {
		callCount++
		return v > 5
	})

	if !ok {
		t.Fatal("FindFirst() should find an element")
	}

	if result != 6 {
		t.Errorf("FindFirst() = %v, want 6", result)
	}

	// Should have checked 1, 2, 3, 4, 5, 6 and stopped
	if callCount > 7 {
		t.Errorf("FindFirst should stop after finding match, called predicate %d times", callCount)
	}
}

// TestFindAny2 tests the FindAny2 function
func TestFindAny2(t *testing.T) {
	tests := []struct {
		name           string
		input          map[string]int
		expectOk       bool
		keyValidator   func(string) bool
		valueValidator func(int) bool
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expectOk: false,
		},
		{
			name:           "single entry",
			input:          map[string]int{"a": 1},
			expectOk:       true,
			keyValidator:   func(k string) bool { return k == "a" },
			valueValidator: func(v int) bool { return v == 1 },
		},
		{
			name:           "multiple entries",
			input:          map[string]int{"a": 1, "b": 2, "c": 3},
			expectOk:       true,
			keyValidator:   func(k string) bool { return k == "a" || k == "b" || k == "c" },
			valueValidator: func(v int) bool { return v >= 1 && v <= 3 },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			key, value, ok := ziter.FindAny2(seq)

			if ok != tt.expectOk {
				t.Errorf("FindAny2() ok = %v, want %v", ok, tt.expectOk)
			}

			if ok {
				if tt.keyValidator != nil && !tt.keyValidator(key) {
					t.Errorf("FindAny2() returned unexpected key = %v", key)
				}
				if tt.valueValidator != nil && !tt.valueValidator(value) {
					t.Errorf("FindAny2() returned unexpected value = %v", value)
				}
			} else {
				if key != "" {
					t.Errorf("FindAny2() should return zero key when ok=false, got %v", key)
				}
				if value != 0 {
					t.Errorf("FindAny2() should return zero value when ok=false, got %v", value)
				}
			}
		})
	}
}

// TestFindAny2EarlyTermination verifies FindAny2 stops after first element
func TestFindAny2EarlyTermination(t *testing.T) {
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

	k, v, ok := ziter.FindAny2(input)

	if !ok {
		t.Fatal("FindAny2() should find an element")
	}

	if k != "a" || v != 1 {
		t.Errorf("FindAny2() = (%v, %v), want (a, 1)", k, v)
	}

	if callCount != 1 {
		t.Errorf("FindAny2 should call yield exactly once, called %d times", callCount)
	}
}

// TestFindFirst2EarlyTermination verifies FindFirst2 stops after finding first match
func TestFindFirst2EarlyTermination(t *testing.T) {
	callCount := 0
	input := func(yield func(string, int) bool) {
		pairs := []struct {
			k string
			v int
		}{
			{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}, {"e", 5},
		}
		for _, p := range pairs {
			if !yield(p.k, p.v) {
				return
			}
		}
	}

	k, v, ok := ziter.FindFirst2(input, func(key string, val int) bool {
		callCount++
		return val > 2
	})

	if !ok {
		t.Fatal("FindFirst2() should find an element")
	}

	if k != "c" || v != 3 {
		t.Errorf("FindFirst2() = (%v, %v), want (c, 3)", k, v)
	}

	// Should have checked (a,1), (b,2), (c,3) and stopped
	if callCount > 4 {
		t.Errorf("FindFirst2 should stop after finding match, called predicate %d times", callCount)
	}
}

// TestFindFirst2 tests the FindFirst2 function
func TestFindFirst2(t *testing.T) {
	tests := []struct {
		name       string
		input      map[string]int
		predicate  func(string, int) bool
		expectedK  string
		expectedV  int
		expectedOk bool
	}{
		{
			name:       "empty map",
			input:      map[string]int{},
			predicate:  func(string, int) bool { return true },
			expectedK:  "",
			expectedV:  0,
			expectedOk: false,
		},
		{
			name:       "no match",
			input:      map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:  func(k string, v int) bool { return v > 10 },
			expectedK:  "",
			expectedV:  0,
			expectedOk: false,
		},
		{
			name:       "match by value",
			input:      map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:  func(_ string, v int) bool { return v == 2 },
			expectedK:  "b",
			expectedV:  2,
			expectedOk: true,
		},
		{
			name:       "match by key",
			input:      map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			predicate:  func(k string, _ int) bool { return k == "banana" },
			expectedK:  "banana",
			expectedV:  2,
			expectedOk: true,
		},
		{
			name:       "match by both key and value",
			input:      map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate:  func(k string, v int) bool { return k >= "c" && v > 2 },
			expectedK:  "", // Any matching key is valid due to undefined map iteration order
			expectedV:  0,
			expectedOk: true,
		},
		{
			name:       "single entry match",
			input:      map[string]int{"x": 10},
			predicate:  func(k string, v int) bool { return k == "x" && v == 10 },
			expectedK:  "x",
			expectedV:  10,
			expectedOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			k, v, ok := ziter.FindFirst2(seq, tt.predicate)

			if ok != tt.expectedOk {
				t.Errorf("FindFirst2() ok = %v, want %v", ok, tt.expectedOk)
			}

			if ok {
				// If expectedK is empty, validate using predicate (for cases with undefined iteration order)
				if tt.expectedK == "" && tt.expectedV == 0 && tt.expectedOk {
					if !tt.predicate(k, v) {
						t.Errorf("FindFirst2() returned (%v, %v) which doesn't match predicate", k, v)
					}
				} else {
					if k != tt.expectedK {
						t.Errorf("FindFirst2() key = %v, want %v", k, tt.expectedK)
					}
					if v != tt.expectedV {
						t.Errorf("FindFirst2() value = %v, want %v", v, tt.expectedV)
					}
				}
			}
		})
	}
}
