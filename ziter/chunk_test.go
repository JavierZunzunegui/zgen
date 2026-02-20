package ziter_test

import (
	"slices"
	"testing"

	"github.com/JavierZunzunegui/zgen/ziter"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  [][]int
	}{
		{
			name:  "empty",
			input: []int{},
			n:     3,
			want:  nil,
		},
		{
			name:  "n zero",
			input: []int{1, 2, 3},
			n:     0,
			want:  nil,
		},
		{
			name:  "n negative",
			input: []int{1, 2, 3},
			n:     -1,
			want:  nil,
		},
		{
			name:  "exact multiple",
			input: []int{1, 2, 3, 4, 5, 6},
			n:     3,
			want:  [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:  "remainder",
			input: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name:  "n equals length",
			input: []int{1, 2, 3},
			n:     3,
			want:  [][]int{{1, 2, 3}},
		},
		{
			name:  "n greater than length",
			input: []int{1, 2},
			n:     5,
			want:  [][]int{{1, 2}},
		},
		{
			name:  "n is 1",
			input: []int{1, 2, 3},
			n:     1,
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "single element",
			input: []int{42},
			n:     3,
			want:  [][]int{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [][]int
			for chunk := range ziter.Chunk(slices.Values(tt.input), tt.n) {
				got = append(got, chunk)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("Chunk() got %d chunks, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if !slices.Equal(got[i], tt.want[i]) {
					t.Errorf("chunk[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestChunkEarlyTermination(t *testing.T) {
	callCount := 0
	seq := func(yield func(int) bool) {
		for i := range 20 {
			callCount++
			if !yield(i) {
				return
			}
		}
	}

	count := 0
	for range ziter.Chunk(seq, 3) {
		count++
		if count == 2 {
			break
		}
	}

	if count != 2 {
		t.Errorf("consumed %d chunks, want 2", count)
	}
	if callCount > 7 {
		t.Errorf("Chunk didn't respect early termination: callCount=%d, expected <= 7", callCount)
	}
}
