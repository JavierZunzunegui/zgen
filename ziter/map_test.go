package ziter_test

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

// TestMap tests the Map function
func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapFunc  func(int) int
		expected []int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			mapFunc:  func(v int) int { return v * 2 },
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{5},
			mapFunc:  func(v int) int { return v * 2 },
			expected: []int{10},
		},
		{
			name:     "double values",
			input:    []int{1, 2, 3, 4, 5},
			mapFunc:  func(v int) int { return v * 2 },
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "identity function",
			input:    []int{1, 2, 3},
			mapFunc:  func(v int) int { return v },
			expected: []int{1, 2, 3},
		},
		{
			name:     "constant function",
			input:    []int{1, 2, 3},
			mapFunc:  func(int) int { return 42 },
			expected: []int{42, 42, 42},
		},
		{
			name:     "negate values",
			input:    []int{1, -2, 3, -4},
			mapFunc:  func(v int) int { return -v },
			expected: []int{-1, 2, -3, 4},
		},
		{
			name:     "square values",
			input:    []int{1, 2, 3, 4},
			mapFunc:  func(v int) int { return v * v },
			expected: []int{1, 4, 9, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			mapped := ziter.Map(seq, tt.mapFunc)
			result := slices.Collect(mapped)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Map() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMapTypeConversion tests Map with type conversions
func TestMapTypeConversion(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		input := slices.Values([]int{1, 2, 3})
		mapped := ziter.Map(input, func(v int) string {
			return strconv.Itoa(v)
		})
		result := slices.Collect(mapped)
		expected := []string{"1", "2", "3"}

		if !slices.Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})

	t.Run("string to length", func(t *testing.T) {
		input := slices.Values([]string{"a", "bb", "ccc"})
		mapped := ziter.Map(input, func(s string) int {
			return len(s)
		})
		result := slices.Collect(mapped)
		expected := []int{1, 2, 3}

		if !slices.Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})
}

// TestMapEarlyTermination verifies Map respects early termination
func TestMapEarlyTermination(t *testing.T) {
	callCount := 0
	input := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	mapped := ziter.Map(input, func(v int) int {
		callCount++
		return v * 2
	})

	// Consume only 3 elements
	count := 0
	for range mapped {
		count++
		if count == 3 {
			break
		}
	}

	if callCount > 4 {
		t.Errorf("Map didn't respect early termination: callCount=%d, expected <= 4", callCount)
	}
}

// TestMap2 tests the Map2 function
func TestMap2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapFunc  func(string, int) (string, int)
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			mapFunc:  func(k string, v int) (string, int) { return k, v },
			expected: map[string]int{},
		},
		{
			name:  "identity",
			input: map[string]int{"a": 1, "b": 2},
			mapFunc: func(k string, v int) (string, int) {
				return k, v
			},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:  "transform both key and value",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			mapFunc: func(k string, v int) (string, int) {
				return k + k, v * 10
			},
			expected: map[string]int{"aa": 10, "bb": 20, "cc": 30},
		},
		{
			name:  "swap key and value types",
			input: map[string]int{"1": 10, "2": 20},
			mapFunc: func(k string, v int) (string, int) {
				return strconv.Itoa(v), func() int { i, _ := strconv.Atoi(k); return i }()
			},
			expected: map[string]int{"10": 1, "20": 2},
		},
		{
			name:  "combine key and value",
			input: map[string]int{"x": 5, "y": 10},
			mapFunc: func(k string, v int) (string, int) {
				return k + strconv.Itoa(v), v + len(k)
			},
			expected: map[string]int{"x5": 6, "y10": 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			mapped := ziter.Map2(seq, tt.mapFunc)
			result := maps.Collect(mapped)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("Map2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMap2TypeConversion tests Map2 with type conversions
func TestMap2TypeConversion(t *testing.T) {
	t.Run("string-int to int-string", func(t *testing.T) {
		input := maps.All(map[string]int{"a": 1, "b": 2})
		mapped := ziter.Map2(input, func(k string, v int) (int, string) {
			return v, k
		})
		result := maps.Collect(mapped)
		expected := map[int]string{1: "a", 2: "b"}

		if !maps.Equal(result, expected) {
			t.Errorf("Map2() = %v, want %v", result, expected)
		}
	})
}

// TestMapKey tests the MapKey function
func TestMapKey(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapFunc  func(string) string
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			mapFunc:  func(k string) string { return k },
			expected: map[string]int{},
		},
		{
			name:     "uppercase keys",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			mapFunc:  func(k string) string { return k + k },
			expected: map[string]int{"aa": 1, "bb": 2, "cc": 3},
		},
		{
			name:     "prefix keys",
			input:    map[string]int{"x": 10, "y": 20},
			mapFunc:  func(k string) string { return "key_" + k },
			expected: map[string]int{"key_x": 10, "key_y": 20},
		},
		{
			name:     "constant key",
			input:    map[string]int{"a": 1, "b": 2},
			mapFunc:  func(k string) string { return "same" },
			expected: map[string]int{}, // Will validate manually due to undefined iteration order
		},
		{
			name:     "values preserved",
			input:    map[string]int{"k1": 100, "k2": 200, "k3": 300},
			mapFunc:  func(k string) string { return k + "_new" },
			expected: map[string]int{"k1_new": 100, "k2_new": 200, "k3_new": 300},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			mapped := ziter.MapKey(seq, tt.mapFunc)
			result := maps.Collect(mapped)

			// Special handling for tests with undefined iteration order
			if tt.name == "constant key" {
				if len(result) != 1 {
					t.Errorf("MapKey() result length = %d, want 1", len(result))
				}
				if v, ok := result["same"]; !ok {
					t.Errorf("MapKey() missing key 'same'")
				} else if v != 1 && v != 2 {
					t.Errorf("MapKey() value = %d, want 1 or 2", v)
				}
			} else if !maps.Equal(result, tt.expected) {
				t.Errorf("MapKey() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMapKeyTypeConversion tests MapKey with type conversion
func TestMapKeyTypeConversion(t *testing.T) {
	t.Run("string key to int key", func(t *testing.T) {
		input := maps.All(map[string]int{"1": 10, "2": 20, "3": 30})
		mapped := ziter.MapKey(input, func(k string) int {
			v, _ := strconv.Atoi(k)
			return v
		})
		result := maps.Collect(mapped)
		expected := map[int]int{1: 10, 2: 20, 3: 30}

		if !maps.Equal(result, expected) {
			t.Errorf("MapKey() = %v, want %v", result, expected)
		}
	})
}

// TestMapKey2 tests the MapKey2 function
func TestMapKey2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapFunc  func(string, int) string
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			mapFunc:  func(k string, v int) string { return k },
			expected: map[string]int{},
		},
		{
			name:  "key depends on value",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			mapFunc: func(k string, v int) string {
				return fmt.Sprintf("%s%d", k, v)
			},
			expected: map[string]int{"a1": 1, "b2": 2, "c3": 3},
		},
		{
			name:  "key is value as string",
			input: map[string]int{"x": 10, "y": 20},
			mapFunc: func(k string, v int) string {
				return strconv.Itoa(v)
			},
			expected: map[string]int{"10": 10, "20": 20},
		},
		{
			name:  "combine key and value",
			input: map[string]int{"foo": 5, "bar": 10},
			mapFunc: func(k string, v int) string {
				if v > 7 {
					return k + "_large"
				}
				return k + "_small"
			},
			expected: map[string]int{"foo_small": 5, "bar_large": 10},
		},
		{
			name:  "values preserved",
			input: map[string]int{"k1": 100, "k2": 200},
			mapFunc: func(k string, v int) string {
				return k + strconv.Itoa(v/10)
			},
			expected: map[string]int{"k110": 100, "k220": 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			mapped := ziter.MapKey2(seq, tt.mapFunc)
			result := maps.Collect(mapped)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("MapKey2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMapValue tests the MapValue function
func TestMapValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapFunc  func(int) int
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			mapFunc:  func(v int) int { return v },
			expected: map[string]int{},
		},
		{
			name:     "double values",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			mapFunc:  func(v int) int { return v * 2 },
			expected: map[string]int{"a": 2, "b": 4, "c": 6},
		},
		{
			name:     "constant value",
			input:    map[string]int{"x": 10, "y": 20, "z": 30},
			mapFunc:  func(v int) int { return 100 },
			expected: map[string]int{"x": 100, "y": 100, "z": 100},
		},
		{
			name:     "negate values",
			input:    map[string]int{"a": 5, "b": -10, "c": 15},
			mapFunc:  func(v int) int { return -v },
			expected: map[string]int{"a": -5, "b": 10, "c": -15},
		},
		{
			name:     "keys preserved",
			input:    map[string]int{"key1": 1, "key2": 2, "key3": 3},
			mapFunc:  func(v int) int { return v * v },
			expected: map[string]int{"key1": 1, "key2": 4, "key3": 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			mapped := ziter.MapValue(seq, tt.mapFunc)
			result := maps.Collect(mapped)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("MapValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMapValueTypeConversion tests MapValue with type conversion
func TestMapValueTypeConversion(t *testing.T) {
	t.Run("int value to string value", func(t *testing.T) {
		input := maps.All(map[string]int{"a": 1, "b": 2, "c": 3})
		mapped := ziter.MapValue(input, func(v int) string {
			return fmt.Sprintf("val%d", v)
		})
		result := maps.Collect(mapped)
		expected := map[string]string{"a": "val1", "b": "val2", "c": "val3"}

		if !maps.Equal(result, expected) {
			t.Errorf("MapValue() = %v, want %v", result, expected)
		}
	})
}

// TestMapValue2 tests the MapValue2 function
func TestMapValue2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapFunc  func(string, int) int
		expected map[string]int
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			mapFunc:  func(k string, v int) int { return v },
			expected: map[string]int{},
		},
		{
			name:  "value depends on key",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			mapFunc: func(k string, v int) int {
				return v + len(k)
			},
			expected: map[string]int{"a": 2, "b": 3, "c": 4},
		},
		{
			name:  "value is key length",
			input: map[string]int{"foo": 10, "x": 20, "hello": 30},
			mapFunc: func(k string, v int) int {
				return len(k)
			},
			expected: map[string]int{"foo": 3, "x": 1, "hello": 5},
		},
		{
			name:  "combine key and value",
			input: map[string]int{"a": 5, "bb": 10, "ccc": 15},
			mapFunc: func(k string, v int) int {
				return v * len(k)
			},
			expected: map[string]int{"a": 5, "bb": 20, "ccc": 45},
		},
		{
			name:  "conditional based on key",
			input: map[string]int{"small": 5, "large": 10},
			mapFunc: func(k string, v int) int {
				if len(k) > 4 {
					return v * 2
				}
				return v
			},
			expected: map[string]int{"small": 10, "large": 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := maps.All(tt.input)
			mapped := ziter.MapValue2(seq, tt.mapFunc)
			result := maps.Collect(mapped)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("MapValue2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestMapValue2TypeConversion tests MapValue2 with type conversion
func TestMapValue2TypeConversion(t *testing.T) {
	t.Run("int value to string value using key", func(t *testing.T) {
		input := maps.All(map[string]int{"a": 1, "b": 2})
		mapped := ziter.MapValue2(input, func(k string, v int) string {
			return k + strconv.Itoa(v)
		})
		result := maps.Collect(mapped)
		expected := map[string]string{"a": "a1", "b": "b2"}

		if !maps.Equal(result, expected) {
			t.Errorf("MapValue2() = %v, want %v", result, expected)
		}
	})
}
