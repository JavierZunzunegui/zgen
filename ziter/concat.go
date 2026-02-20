package ziter

import "iter"

// Concat returns a sequence that yields all elements from each of seqs in order.
func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}

// Concat2 returns a sequence that yields all key-value pairs from each of seqs in order.
func Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}
