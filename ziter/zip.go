package ziter

import "iter"

// Zip combines two sequences element-wise into a single Seq2.
// It stops when either sequence is exhausted.
func Zip[A, B any](seqA iter.Seq[A], seqB iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		nextB, stopB := iter.Pull(seqB)
		defer stopB()
		seqA(func(a A) bool {
			b, ok := nextB()
			if !ok {
				return false
			}
			return yield(a, b)
		})
	}
}
