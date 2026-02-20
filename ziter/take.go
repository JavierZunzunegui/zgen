package ziter

import "iter"

// Take returns a sequence of up to n elements from the start of seq.
func Take[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		if n <= 0 {
			return
		}
		i := 0
		seq(func(v V) bool {
			if !yield(v) {
				return false
			}
			i++
			return i < n
		})
	}
}

// Take2 is like [Take] but with [iter.Seq2].
func Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if n <= 0 {
			return
		}
		i := 0
		seq(func(k K, v V) bool {
			if !yield(k, v) {
				return false
			}
			i++
			return i < n
		})
	}
}

// Drop returns a sequence that skips the first n elements from seq and yields the rest.
func Drop[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		seq(func(v V) bool {
			if i < n {
				i++
				return true
			}
			return yield(v)
		})
	}
}

// Drop2 is like [Drop] but with [iter.Seq2].
func Drop2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		seq(func(k K, v V) bool {
			if i < n {
				i++
				return true
			}
			return yield(k, v)
		})
	}
}
