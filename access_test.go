package zgen_test

import (
	"errors"
	"testing"

	"github.com/JavierZunzunegui/zgen"
)

func TestBoolLast2(t *testing.T) {
	tests := []struct {
		name string
		arg  bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := zgen.BoolLast2("ignored", tc.arg)
			if got != tc.arg {
				t.Errorf("got %v, want %v", got, tc.arg)
			}
		})
	}
}

func TestErrLast2(t *testing.T) {
	sentinel := errors.New("sentinel")
	tests := []struct {
		name string
		arg  error
	}{
		{"nil", nil},
		{"non-nil", sentinel},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := zgen.ErrLast2("ignored", tc.arg)
			if got != tc.arg {
				t.Errorf("got %v, want %v", got, tc.arg)
			}
		})
	}
}
