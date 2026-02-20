package ziter

import "iter"

// Dedup removes duplicate values from seq, yielding each unique value only once.
// V must be comparable.
func Dedup[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[V]struct{})
		seq(func(v V) bool {
			if _, ok := seen[v]; ok {
				return true
			}
			seen[v] = struct{}{}
			return yield(v)
		})
	}
}

// Dedup2 is like [Dedup] but for [iter.Seq2], deduplicating on the K-V pair jointly.
// Both K and V must be comparable.
func Dedup2[K, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		k K
		v V
	}
	return func(yield func(K, V) bool) {
		seen := make(map[pair]struct{})
		seq(func(k K, v V) bool {
			p := pair{k, v}
			if _, ok := seen[p]; ok {
				return true
			}
			seen[p] = struct{}{}
			return yield(k, v)
		})
	}
}
