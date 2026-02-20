package ziter

import "iter"

// Chunk groups the elements of seq into slices of up to n elements each.
// The last chunk may have fewer than n elements.
// If n <= 0 or seq is empty, the returned sequence is empty.
func Chunk[V any](seq iter.Seq[V], n int) iter.Seq[[]V] {
	return func(yield func([]V) bool) {
		if n <= 0 {
			return
		}
		chunk := make([]V, 0, n)
		stopped := false
		seq(func(v V) bool {
			chunk = append(chunk, v)
			if len(chunk) == n {
				if !yield(chunk) {
					stopped = true
					return false
				}
				chunk = make([]V, 0, n)
			}
			return true
		})
		if len(chunk) > 0 && !stopped {
			yield(chunk)
		}
	}
}
