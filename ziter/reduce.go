package ziter

import (
	"cmp"
	"iter"
)

// Reduce accumulates the elements of seq into a single value using f.
// The first element is used as the initial accumulator.
// If seq is empty, it returns the zero value and false.
func Reduce[V any](seq iter.Seq[V], f func(V, V) V) (out V, ok bool) {
	seq(func(v V) bool {
		if !ok {
			out, ok = v, true
		} else {
			out = f(out, v)
		}
		return true
	})
	return out, ok
}

// Reduce2 is like [Reduce] but with [iter.Seq2].
// The first key-value pair is used as the initial accumulator.
// If seq is empty, it returns zero values and false.
func Reduce2[K, V any](seq iter.Seq2[K, V], f func(K, V, K, V) (K, V)) (outK K, outV V, ok bool) {
	seq(func(k K, v V) bool {
		if !ok {
			outK, outV, ok = k, v, true
		} else {
			outK, outV = f(outK, outV, k, v)
		}
		return true
	})
	return outK, outV, ok
}

// Aggregate accumulates the elements of seq into a single value of a
// (potentially different) type using f, starting from init.
func Aggregate[V, A any](seq iter.Seq[V], init A, f func(A, V) A) A {
	seq(func(v V) bool {
		init = f(init, v)
		return true
	})
	return init
}

// Aggregate2 is like [Aggregate] but with [iter.Seq2].
func Aggregate2[K, V, A any](seq iter.Seq2[K, V], init A, f func(A, K, V) A) A {
	seq(func(k K, v V) bool {
		init = f(init, k, v)
		return true
	})
	return init
}

// Count returns the number of elements in seq.
func Count[V any](seq iter.Seq[V]) int {
	var count int
	seq(func(V) bool {
		count++
		return true
	})
	return count
}

// Count2 returns the number of key-value pairs in seq.
func Count2[K, V any](seq iter.Seq2[K, V]) int {
	var count int
	seq(func(K, V) bool {
		count++
		return true
	})
	return count
}

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
