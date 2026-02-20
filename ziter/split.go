package ziter

// Note that splitting sequences is not done particularly efficiently.

import (
	"iter"
	"slices"

	"github.com/JavierZunzunegui/zgen"
)

// Split partitions seq into two sequences: the first yields elements that
// satisfy f, the second yields elements that do not. The input sequence is
// collected eagerly so both outputs can be iterated independently.
func Split[V any](seq iter.Seq[V], f func(V) bool) (iter.Seq[V], iter.Seq[V]) {
	all := slices.Collect(seq)
	negate := func(v V) bool { return !f(v) }
	return Filter(slices.Values(all), f), Filter(slices.Values(all), negate)
}

// Split2 partitions seq into two sequences: the first yields pairs that
// satisfy f, the second yields pairs that do not. The input sequence is
// collected eagerly so both outputs can be iterated independently.
func Split2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	left, right := Split(
		ToSeq1(seq, zgen.NewPair),
		func(p zgen.Pair[K, V]) bool { return f(p.Both()) },
	)

	return ToSeq2(left, zgen.Pair[K, V].Both), ToSeq2(right, zgen.Pair[K, V].Both)
}

// SplitKey partitions seq by applying f to each key, equivalent to Split2
// with a predicate that ignores the value.
func SplitKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	return Split2(seq, func(k K, v V) bool { return f(k) })
}

// SplitValue partitions seq by applying f to each value, equivalent to Split2
// with a predicate that ignores the key.
func SplitValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	return Split2(seq, func(k K, v V) bool { return f(v) })
}
