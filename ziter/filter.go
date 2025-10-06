package ziter

import "iter"

// Filter removes entries that are false on the given function from the resulting sequence.
func Filter[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(v V) bool {
			if f(v) {
				return yield(v)
			}
			return true
		})
	}
}

// Filter2 is like [Filter] but with [iter.Seq2].
func Filter2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if f(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}

// FilterKey is like [Filter2] but only filters based on the key.
func FilterKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if f(k) {
				return yield(k, v)
			}
			return true
		})
	}
}

// FilterValue is like [Filter2] but only filters based on the value.
func FilterValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if f(v) {
				return yield(k, v)
			}
			return true
		})
	}
}
