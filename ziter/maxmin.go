package ziter

import (
	"cmp"
	"iter"
)

// MaxFunc returns the maximum element of seq and true, using cmp to compare values.
// If seq is empty, it returns the zero value and false.
// cmp(a, b) must return a negative int if a < b, zero if a == b, a positive int if a > b.
func MaxFunc[V any](seq iter.Seq[V], cmp func(V, V) int) (out V, ok bool) {
	seq(func(v V) bool {
		if !ok || cmp(v, out) > 0 {
			out, ok = v, true
		}
		return true
	})
	return out, ok
}

// MinFunc returns the minimum element of seq and true, using cmp to compare values.
// If seq is empty, it returns the zero value and false.
// cmp(a, b) must return a negative int if a < b, zero if a == b, a positive int if a > b.
func MinFunc[V any](seq iter.Seq[V], cmp func(V, V) int) (out V, ok bool) {
	seq(func(v V) bool {
		if !ok || cmp(v, out) < 0 {
			out, ok = v, true
		}
		return true
	})
	return out, ok
}

// Max returns the maximum element of seq and true.
// If seq is empty, it returns the zero value and false.
func Max[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	return MaxFunc(seq, cmp.Compare)
}

// Min returns the minimum element of seq and true.
// If seq is empty, it returns the zero value and false.
func Min[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	return MinFunc(seq, cmp.Compare)
}
