package zgen_test

import (
	"errors"
	"testing"

	"github.com/JavierZunzunegui/zgen"
)

func TestPtrTo(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		p := zgen.PtrTo(42)
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != 42 {
			t.Errorf("got %v, want 42", *p)
		}
	})

	t.Run("string", func(t *testing.T) {
		p := zgen.PtrTo("hello")
		if p == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *p != "hello" {
			t.Errorf("got %v, want hello", *p)
		}
	})
}

func TestIsType(t *testing.T) {
	t.Run("concrete match", func(t *testing.T) {
		if !zgen.IsType[int](42) {
			t.Error("expected true for int(42)")
		}
	})

	t.Run("concrete mismatch", func(t *testing.T) {
		if zgen.IsType[int]("hello") {
			t.Error("expected false for string as int")
		}
	})

	t.Run("interface match", func(t *testing.T) {
		e := errors.New("e")
		if !zgen.IsType[error](e) {
			t.Error("expected true for error interface")
		}
	})

	t.Run("nil", func(t *testing.T) {
		if zgen.IsType[int](nil) {
			t.Error("expected false for nil as int")
		}
	})
}

func TestCast(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v, ok := zgen.Cast[int](42)
		if !ok || v != 42 {
			t.Errorf("got (%v, %v), want (42, true)", v, ok)
		}
	})

	t.Run("failure", func(t *testing.T) {
		v, ok := zgen.Cast[int]("hello")
		if ok || v != 0 {
			t.Errorf("got (%v, %v), want (0, false)", v, ok)
		}
	})
}
