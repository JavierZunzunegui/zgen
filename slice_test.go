package zgen_test

import (
	"testing"

	"github.com/JavierZunzunegui/zgen"
)

func TestNoCap(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var s []int
		got := zgen.NoCap(s)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
		if cap(got) != 0 {
			t.Errorf("expected cap 0, got %d", cap(got))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		s := make([]int, 0)
		got := zgen.NoCap(s)
		if cap(got) != 0 {
			t.Errorf("expected cap 0, got %d", cap(got))
		}
	})

	t.Run("len equals cap", func(t *testing.T) {
		s := []int{1, 2, 3}
		got := zgen.NoCap(s)
		if cap(got) != len(got) {
			t.Errorf("cap %d != len %d", cap(got), len(got))
		}
		if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
			t.Errorf("contents changed: %v", got)
		}
	})

	t.Run("excess cap trimmed", func(t *testing.T) {
		base := make([]int, 10)
		base[0], base[1], base[2] = 7, 8, 9
		s := base[:3:10]
		got := zgen.NoCap(s)
		if cap(got) != len(got) {
			t.Errorf("cap %d != len %d", cap(got), len(got))
		}
		if len(got) != 3 || got[0] != 7 || got[1] != 8 || got[2] != 9 {
			t.Errorf("contents changed: %v", got)
		}
	})
}
