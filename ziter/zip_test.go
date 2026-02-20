package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

func TestZip(t *testing.T) {
	tests := []struct {
		name      string
		a         []int
		b         []string
		expectedA []int
		expectedB []string
	}{
		{
			name:      "both empty",
			a:         []int{},
			b:         []string{},
			expectedA: []int{},
			expectedB: []string{},
		},
		{
			name:      "first empty",
			a:         []int{},
			b:         []string{"x", "y"},
			expectedA: []int{},
			expectedB: []string{},
		},
		{
			name:      "second empty",
			a:         []int{1, 2},
			b:         []string{},
			expectedA: []int{},
			expectedB: []string{},
		},
		{
			name:      "equal length",
			a:         []int{1, 2, 3},
			b:         []string{"a", "b", "c"},
			expectedA: []int{1, 2, 3},
			expectedB: []string{"a", "b", "c"},
		},
		{
			name:      "first shorter",
			a:         []int{1, 2},
			b:         []string{"a", "b", "c"},
			expectedA: []int{1, 2},
			expectedB: []string{"a", "b"},
		},
		{
			name:      "second shorter",
			a:         []int{1, 2, 3},
			b:         []string{"a", "b"},
			expectedA: []int{1, 2},
			expectedB: []string{"a", "b"},
		},
		{
			name:      "single element each",
			a:         []int{42},
			b:         []string{"only"},
			expectedA: []int{42},
			expectedB: []string{"only"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotA []int
			var gotB []string
			for a, b := range ziter.Zip(slices.Values(tt.a), slices.Values(tt.b)) {
				gotA = append(gotA, a)
				gotB = append(gotB, b)
			}
			if !slices.Equal(gotA, tt.expectedA) {
				t.Errorf("keys: got %v, want %v", gotA, tt.expectedA)
			}
			if !slices.Equal(gotB, tt.expectedB) {
				t.Errorf("values: got %v, want %v", gotB, tt.expectedB)
			}
		})
	}
}

func TestZipEarlyTermination(t *testing.T) {
	var gotA []int
	var gotB []string
	for a, b := range ziter.Zip(slices.Values([]int{1, 2, 3, 4, 5}), slices.Values([]string{"a", "b", "c", "d", "e"})) {
		gotA = append(gotA, a)
		gotB = append(gotB, b)
		if len(gotA) == 2 {
			break
		}
	}
	if !slices.Equal(gotA, []int{1, 2}) {
		t.Errorf("keys: got %v, want [1 2]", gotA)
	}
	if !slices.Equal(gotB, []string{"a", "b"}) {
		t.Errorf("values: got %v, want [a b]", gotB)
	}
}
