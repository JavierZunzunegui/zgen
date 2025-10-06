package ziter

import "iter"

// Map changes entries based on the given function from the resulting sequence.
// No entries are added or removed.
func Map[V1, V2 any](seq iter.Seq[V1], f func(V1) V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		seq(func(v V1) bool {
			return yield(f(v))
		})
	}
}

// Map2 is like [Map] but with [iter.Seq2].
func Map2[K1, V1, K2, V2 any](seq iter.Seq2[K1, V1], f func(K1, V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		seq(func(k K1, v V1) bool {
			return yield(f(k, v))
		})
	}
}

// MapKey is like [Map2] but only transforms the key.
func MapKey[K1, V, K2 any](seq iter.Seq2[K1, V], f func(K1) K2) iter.Seq2[K2, V] {
	return func(yield func(K2, V) bool) {
		seq(func(k K1, v V) bool {
			return yield(f(k), v)
		})
	}
}

// MapValue is like [Map2] but only transforms the value.
func MapValue[K, V1, V2 any](seq iter.Seq2[K, V1], f func(V1) V2) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		seq(func(k K, v V1) bool {
			return yield(k, f(v))
		})
	}
}
