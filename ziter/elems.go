package ziter

import "iter"

func Exists[V any](seq iter.Seq[V]) bool {
	_, ok := FindAny(seq)
	return ok
}

func Count[V any](seq iter.Seq[V]) int {
	var count int
	seq(func(V) bool {
		count++
		return true
	})
	return count
}

func Count2[K, V any](seq iter.Seq2[K, V]) int {
	var count int
	seq(func(K, V) bool {
		count++
		return true
	})
	return count
}
