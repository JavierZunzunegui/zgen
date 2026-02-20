package ziter

import "iter"

// Exists reports whether seq contains at least one element.
func Exists[V any](seq iter.Seq[V]) bool {
	_, ok := FindAny(seq)
	return ok
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
