package ziter

import (
	"iter"

	"github.com/JavierZunzunegui/zgen"
)

func FindAny[V any](seq iter.Seq[V]) (out V, ok bool) {
	seq(func(v V) bool {
		out, ok = v, true
		return false
	})
	return out, ok
}

func FindFirst[V any](seq iter.Seq[V], f func(V) bool) (V, bool) {
	return FindAny(Filter(seq, f))
}

func FindAny2[K, V any](seq iter.Seq2[K, V]) (outK K, outV V, ok bool) {
	seq(func(k K, v V) bool {
		outK, outV, ok = k, v, true
		return false
	})
	return outK, outV, ok
}

func FindFirst2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (K, V, bool) {
	return FindAny2(Filter2(seq, f))
}

func Exists[V any](seq iter.Seq[V]) bool {
	return zgen.BoolLast2(FindAny(seq))
}
