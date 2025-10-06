package ziter

import "iter"

// Flatten TODO
func Flatten[V1, V2 any](seq iter.Seq[V1], f func(V1) []V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		seq(func(v V1) bool {
			for _, v2 := range f(v) {
				if !yield(v2) {
					return false
				}
			}
			return true
		})
	}
}

// FlattenKeys TODO
func FlattenKeys[K1, V, K2 any](seq iter.Seq2[K1, V], f func(K1) []K2) iter.Seq2[K2, V] {
	return func(yield func(K2, V) bool) {
		seq(func(k K1, v V) bool {
			for _, k2 := range f(k) {
				if !yield(k2, v) {
					return false
				}
			}
			return true
		})
	}
}

// FlattenValues TODO
func FlattenValues[K, V1, V2 any](seq iter.Seq2[K, V1], f func(V1) []V2) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		seq(func(k K, v V1) bool {
			for _, v2 := range f(v) {
				if !yield(k, v2) {
					return false
				}
			}
			return true
		})
	}
}
