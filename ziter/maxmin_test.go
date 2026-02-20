package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

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
