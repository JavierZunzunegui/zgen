package ziter

import "iter"

// FindAny TODO
func FindAny[V any](seq iter.Seq[V]) (out V, ok bool) {
	seq(func(v V) bool {
		out, ok = v, true
		return false
	})
	return out, ok
}

// FindFirst TODO
func FindFirst[V any](seq iter.Seq[V], f func(V) bool) (V, bool) {
	return FindAny(Filter(seq, f))
}

// FindAny2 TODO
func FindAny2[K, V any](seq iter.Seq2[K, V]) (outK K, outV V, ok bool) {
	seq(func(k K, v V) bool {
		outK, outV, ok = k, v, true
		return false
	})
	return outK, outV, ok
}

// FindFirst2 TODO
func FindFirst2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (K, V, bool) {
	return FindAny2(Filter2(seq, f))
}
