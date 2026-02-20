package ziter

import (
	"iter"
)

// FindAny returns an arbitrary element from seq and true, or the zero value
// and false if seq is empty. It consumes exactly one element.
func FindAny[V any](seq iter.Seq[V]) (out V, ok bool) {
	seq(func(v V) bool {
		out, ok = v, true
		return false
	})
	return out, ok
}

// FindFirst returns the first element of seq that satisfies f and true, or
// the zero value and false if no such element exists.
func FindFirst[V any](seq iter.Seq[V], f func(V) bool) (V, bool) {
	return FindAny(Filter(seq, f))
}

// FindAny2 returns an arbitrary key-value pair from seq and true, or zero
// values and false if seq is empty. It consumes exactly one element.
func FindAny2[K, V any](seq iter.Seq2[K, V]) (outK K, outV V, ok bool) {
	seq(func(k K, v V) bool {
		outK, outV, ok = k, v, true
		return false
	})
	return outK, outV, ok
}

// FindFirst2 returns the first key-value pair of seq that satisfies f and true,
// or zero values and false if no such pair exists.
func FindFirst2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (K, V, bool) {
	return FindAny2(Filter2(seq, f))
}

// Exists reports whether seq contains at least one element.
func Exists[V any](seq iter.Seq[V]) bool {
	_, ok := FindAny(seq)
	return ok
}
