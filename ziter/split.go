package ziter

// Note that splitting sequences is not done particularly efficiently.

import (
	"iter"
	"slices"
)

// Split TODO.
func Split[V any](seq iter.Seq[V], f func(V) bool) (iter.Seq[V], iter.Seq[V]) {
	all := slices.Collect(seq)
	negate := func(v V) bool { return !f(v) }
	return Filter(slices.Values(all), f), Filter(slices.Values(all), negate)
}

// Split2 TODO
func Split2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	type pair = struct {
		k K
		v V
	}

	left, right := Split(
		ToSeq1(seq, func(k K, v V) pair { return pair{k, v} }),
		func(p pair) bool { return f(p.k, p.v) },
	)

	toKV := func(p pair) (K, V) { return p.k, p.v }
	return ToSeq2(left, toKV), ToSeq2(right, toKV)
}

// SplitKey TODO
func SplitKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	return Split2(seq, func(k K, v V) bool { return f(k) })
}

// SplitValue TODO
func SplitValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	return Split2(seq, func(k K, v V) bool { return f(v) })
}
