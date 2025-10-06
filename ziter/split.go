package ziter

import "iter"

// Split TODO.
func Split[V any](seq iter.Seq[V], f func(V) bool) (iter.Seq[V], iter.Seq[V]) {
	negate := func(v V) bool { return !f(v) }
	return Filter(seq, f), Filter(seq, negate)
}

// Split2 TODO
func Split2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	negate := func(k K, v V) bool { return !f(k, v) }
	return Filter2(seq, f), Filter2(seq, negate)
}

// SplitKey TODO
func SplitKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	negate := func(k K) bool { return !f(k) }
	return FilterKey(seq, f), FilterKey(seq, negate)
}

// SplitValue TODO
func SplitValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	negate := func(v V) bool { return !f(v) }
	return FilterValue(seq, f), FilterValue(seq, negate)
}
