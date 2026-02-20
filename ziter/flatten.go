package ziter

import "iter"

// Flatten returns a sequence that applies f to each element of seq and yields
// the resulting slice elements in order.
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

// FlattenKeys returns a sequence that applies f to each key of seq and yields
// one (k2, v) pair per element of the resulting slice, pairing each new key
// with the original value.
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

// FlattenValues returns a sequence that applies f to each value of seq and yields
// one (k, v2) pair per element of the resulting slice, pairing each new value
// with the original key.
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
