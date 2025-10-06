package ziter

import "iter"

// KeyBy converts a [iter.Seq] to a [iter.Seq2] maintaing the prior entries as values.
func KeyBy[V, K2 any](seq iter.Seq[V], f func(V) K2) iter.Seq2[K2, V] {
	return func(yield func(K2, V) bool) {
		seq(func(v V) bool {
			return yield(f(v), v)
		})
	}
}

// ValueBy converts a [iter.Seq] to a [iter.Seq2] maintaing the prior entries as keys.
func ValueBy[V, V2 any](seq iter.Seq[V], f func(V) V2) iter.Seq2[V, V2] {
	return func(yield func(V, V2) bool) {
		seq(func(v V) bool {
			return yield(v, f(v))
		})
	}
}

// Keys converts a [iter.Seq2] to a [iter.Seq1] maintaing the prior keys as entries.
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		seq(func(k K, v V) bool {
			return yield(k)
		})
	}
}

// Values converts a [iter.Seq2] to a [iter.Seq1] maintaing the prior values as entries.
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(k K, v V) bool {
			return yield(v)
		})
	}
}
