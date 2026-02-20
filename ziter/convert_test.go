package ziter_test

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestToSeq2 tests the ToSeq2 function
func TestToSeq2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		seq := slices.Values([]string{})
		seq2 := ziter.ToSeq2(seq, func(s string) (int, string) { return len(s), s })
		result := maps.Collect(seq2)
		if len(result) != 0 {
			t.Errorf("ToSeq2() = %v, want empty", result)
		}
	})

	t.Run("string to (length, value)", func(t *testing.T) {
		seq := slices.Values([]string{"a", "bb", "ccc"})
		seq2 := ziter.ToSeq2(seq, func(s string) (int, string) { return len(s), s })
		result := maps.Collect(seq2)
		expected := map[int]string{1: "a", 2: "bb", 3: "ccc"}
		if !maps.Equal(result, expected) {
			t.Errorf("ToSeq2() = %v, want %v", result, expected)
		}
	})
}

// TestToSeq1 tests the ToSeq1 function
func TestToSeq1(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		seq2 := maps.All(map[string]int{})
		seq := ziter.ToSeq1(seq2, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) })
		result := slices.Collect(seq)
		if len(result) != 0 {
			t.Errorf("ToSeq1() = %v, want empty", result)
		}
	})

	t.Run("single entry", func(t *testing.T) {
		// Use a deterministic Seq2 (not a map) to avoid ordering issues.
		seq2 := func(yield func(string, int) bool) {
			yield("x", 7)
		}
		seq := ziter.ToSeq1(seq2, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) })
		result := slices.Collect(seq)
		expected := []string{"x:7"}
		if !slices.Equal(result, expected) {
			t.Errorf("ToSeq1() = %v, want %v", result, expected)
		}
	})

	t.Run("multiple entries preserve order", func(t *testing.T) {
		seq2 := func(yield func(string, int) bool) {
			for _, p := range []struct {
				k string
				v int
			}{{"a", 1}, {"b", 2}, {"c", 3}} {
				if !yield(p.k, p.v) {
					return
				}
			}
		}
		seq := ziter.ToSeq1(seq2, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) })
		result := slices.Collect(seq)
		expected := []string{"a:1", "b:2", "c:3"}
		if !slices.Equal(result, expected) {
			t.Errorf("ToSeq1() = %v, want %v", result, expected)
		}
	})
}

// TestKeys tests the Keys function
func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		validate func([]string) bool // validates all expected keys are present
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			validate: func(keys []string) bool { return len(keys) == 0 },
		},
		{
			name:  "single entry",
			input: map[string]int{"a": 1},
			validate: func(keys []string) bool {
				return len(keys) == 1 && keys[0] == "a"
			},
		},
		{
			name:  "multiple entries",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			validate: func(keys []string) bool {
				if len(keys) != 3 {
					return false
				}
				return slices.Contains(keys, "a") &&
					slices.Contains(keys, "b") &&
					slices.Contains(keys, "c")
			},
		},
		{
			name:  "duplicate values different keys",
			input: map[string]int{"x": 10, "y": 10, "z": 10},
			validate: func(keys []string) bool {
				if len(keys) != 3 {
					return false
				}
				return slices.Contains(keys, "x") &&
					slices.Contains(keys, "y") &&
					slices.Contains(keys, "z")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			keys := ziter.Keys(seq)
			result := slices.Collect(keys)

			if !tt.validate(result) {
				t.Errorf("Keys() = %v, validation failed", result)
			}
		})
	}
}

// TestValues tests the Values function
func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		validate func([]int) bool // validates all expected values are present
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			validate: func(vals []int) bool { return len(vals) == 0 },
		},
		{
			name:  "single entry",
			input: map[string]int{"a": 1},
			validate: func(vals []int) bool {
				return len(vals) == 1 && vals[0] == 1
			},
		},
		{
			name:  "multiple entries",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			validate: func(vals []int) bool {
				if len(vals) != 3 {
					return false
				}
				return slices.Contains(vals, 1) &&
					slices.Contains(vals, 2) &&
					slices.Contains(vals, 3)
			},
		},
		{
			name:  "duplicate values",
			input: map[string]int{"x": 5, "y": 5, "z": 5},
			validate: func(vals []int) bool {
				if len(vals) != 3 {
					return false
				}
				for _, v := range vals {
					if v != 5 {
						return false
					}
				}
				return true
			},
		},
		{
			name:  "mixed values",
			input: map[string]int{"neg": -5, "zero": 0, "pos": 10},
			validate: func(vals []int) bool {
				if len(vals) != 3 {
					return false
				}
				return slices.Contains(vals, -5) &&
					slices.Contains(vals, 0) &&
					slices.Contains(vals, 10)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			values := ziter.Values(seq)
			result := slices.Collect(values)

			if !tt.validate(result) {
				t.Errorf("Values() = %v, validation failed", result)
			}
		})
	}
}

// TestKeyBy tests the KeyBy function
func TestKeyBy(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		keyFunc  func(string) int
		expected map[int]string
	}{
		{
			name:     "empty sequence",
			input:    []string{},
			keyFunc:  func(s string) int { return len(s) },
			expected: map[int]string{},
		},
		{
			name:     "single element",
			input:    []string{"hello"},
			keyFunc:  func(s string) int { return len(s) },
			expected: map[int]string{5: "hello"},
		},
		{
			name:     "multiple elements by length",
			input:    []string{"a", "bb", "ccc"},
			keyFunc:  func(s string) int { return len(s) },
			expected: map[int]string{1: "a", 2: "bb", 3: "ccc"},
		},
		{
			name:  "index as key",
			input: []string{"first", "second", "third"},
			keyFunc: func() func(string) int {
				idx := -1
				return func(string) int {
					idx++
					return idx
				}
			}(),
			expected: map[int]string{0: "first", 1: "second", 2: "third"},
		},
		{
			name:     "constant key function",
			input:    []string{"a", "b", "c"},
			keyFunc:  func(string) int { return 42 },
			expected: map[int]string{42: "c"}, // last one wins
		},
		{
			name:     "numeric strings to int",
			input:    []string{"10", "20", "30"},
			keyFunc:  func(s string) int { v, _ := strconv.Atoi(s); return v },
			expected: map[int]string{10: "10", 20: "20", 30: "30"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			keyed := ziter.KeyBy(seq, tt.keyFunc)
			result := maps.Collect(keyed)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("KeyBy() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestKeyByTypes tests KeyBy with different type transformations
func TestKeyByTypes(t *testing.T) {
	t.Run("int to string key", func(t *testing.T) {
		input := slices.Values([]int{1, 2, 3})
		keyed := ziter.KeyBy(input, func(v int) string {
			return fmt.Sprintf("key%d", v)
		})
		result := maps.Collect(keyed)
		expected := map[string]int{"key1": 1, "key2": 2, "key3": 3}

		if !maps.Equal(result, expected) {
			t.Errorf("KeyBy() = %v, want %v", result, expected)
		}
	})

	t.Run("string to first char key", func(t *testing.T) {
		input := slices.Values([]string{"apple", "banana", "cherry"})
		keyed := ziter.KeyBy(input, func(s string) byte {
			if len(s) > 0 {
				return s[0]
			}
			return 0
		})
		result := maps.Collect(keyed)
		expected := map[byte]string{'a': "apple", 'b': "banana", 'c': "cherry"}

		if !maps.Equal(result, expected) {
			t.Errorf("KeyBy() = %v, want %v", result, expected)
		}
	})
}

// TestValueBy tests the ValueBy function
func TestValueBy(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		valueFunc func(string) int
		expected  map[string]int
	}{
		{
			name:      "empty sequence",
			input:     []string{},
			valueFunc: func(s string) int { return len(s) },
			expected:  map[string]int{},
		},
		{
			name:      "single element",
			input:     []string{"hello"},
			valueFunc: func(s string) int { return len(s) },
			expected:  map[string]int{"hello": 5},
		},
		{
			name:      "length as value",
			input:     []string{"a", "bb", "ccc"},
			valueFunc: func(s string) int { return len(s) },
			expected:  map[string]int{"a": 1, "bb": 2, "ccc": 3},
		},
		{
			name:      "constant value",
			input:     []string{"x", "y", "z"},
			valueFunc: func(string) int { return 100 },
			expected:  map[string]int{"x": 100, "y": 100, "z": 100},
		},
		{
			name:      "duplicate keys",
			input:     []string{"key", "key", "key"},
			valueFunc: func(s string) int { return len(s) },
			expected:  map[string]int{"key": 3}, // last wins
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			valued := ziter.ValueBy(seq, tt.valueFunc)
			result := maps.Collect(valued)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("ValueBy() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestValueByTypes tests ValueBy with different type transformations
func TestValueByTypes(t *testing.T) {
	t.Run("int to string value", func(t *testing.T) {
		input := slices.Values([]int{1, 2, 3})
		valued := ziter.ValueBy(input, func(v int) string {
			return fmt.Sprintf("val%d", v*10)
		})
		result := maps.Collect(valued)
		expected := map[int]string{1: "val10", 2: "val20", 3: "val30"}

		if !maps.Equal(result, expected) {
			t.Errorf("ValueBy() = %v, want %v", result, expected)
		}
	})

	t.Run("string to uppercase value", func(t *testing.T) {
		input := slices.Values([]string{"a", "b", "c"})
		valued := ziter.ValueBy(input, func(s string) string {
			return s + s
		})
		result := maps.Collect(valued)
		expected := map[string]string{"a": "aa", "b": "bb", "c": "cc"}

		if !maps.Equal(result, expected) {
			t.Errorf("ValueBy() = %v, want %v", result, expected)
		}
	})
}

func TestSingle(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		got := slices.Collect(ziter.Single(42))
		if !slices.Equal(got, []int{42}) {
			t.Errorf("Single(42) = %v, want [42]", got)
		}
	})

	t.Run("string", func(t *testing.T) {
		got := slices.Collect(ziter.Single("hello"))
		if !slices.Equal(got, []string{"hello"}) {
			t.Errorf("Single(\"hello\") = %v, want [hello]", got)
		}
	})

	t.Run("early termination", func(t *testing.T) {
		count := 0
		for range ziter.Single(1) {
			count++
			break
		}
		if count != 1 {
			t.Errorf("count = %d, want 1", count)
		}
	})
}

func TestSingle2(t *testing.T) {
	t.Run("int string", func(t *testing.T) {
		got := maps.Collect(ziter.Single2(1, "one"))
		expected := map[int]string{1: "one"}
		if !maps.Equal(got, expected) {
			t.Errorf("Single2(1, \"one\") = %v, want %v", got, expected)
		}
	})

	t.Run("string int", func(t *testing.T) {
		got := maps.Collect(ziter.Single2("key", 99))
		expected := map[string]int{"key": 99}
		if !maps.Equal(got, expected) {
			t.Errorf("Single2(\"key\", 99) = %v, want %v", got, expected)
		}
	})

	t.Run("early termination", func(t *testing.T) {
		count := 0
		for range ziter.Single2("a", 1) {
			count++
			break
		}
		if count != 1 {
			t.Errorf("count = %d, want 1", count)
		}
	})
}

// TestKeysEarlyTermination verifies Keys respects early termination
func TestKeysEarlyTermination(t *testing.T) {
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

	keys := ziter.Keys(input)

	// Consume only 2 keys
	count := 0
	for range keys {
		count++
		if count == 2 {
			break
		}
	}

	if callCount > 3 {
		t.Errorf("Keys didn't respect early termination: callCount=%d, expected <= 3", callCount)
	}
}

// TestValuesEarlyTermination verifies Values respects early termination
func TestValuesEarlyTermination(t *testing.T) {
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

	values := ziter.Values(input)

	// Consume only 2 values
	count := 0
	for range values {
		count++
		if count == 2 {
			break
		}
	}

	if callCount > 3 {
		t.Errorf("Values didn't respect early termination: callCount=%d, expected <= 3", callCount)
	}
}

// TestKeyByEarlyTermination verifies KeyBy respects early termination
func TestKeyByEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	keyed := ziter.KeyBy(input, func(v int) string {
		callCount++
		return fmt.Sprintf("key%d", v)
	})

	// Consume only 3 elements
	count := 0
	for range keyed {
		count++
		if count == 3 {
			break
		}
	}

	if callCount > 4 {
		t.Errorf("KeyBy didn't respect early termination: callCount=%d, expected <= 4", callCount)
	}
}

// TestValueByEarlyTermination verifies ValueBy respects early termination
func TestValueByEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	valued := ziter.ValueBy(input, func(v int) string {
		callCount++
		return fmt.Sprintf("val%d", v)
	})

	// Consume only 3 elements
	count := 0
	for range valued {
		count++
		if count == 3 {
			break
		}
	}

	if callCount > 4 {
		t.Errorf("ValueBy didn't respect early termination: callCount=%d, expected <= 4", callCount)
	}
}
