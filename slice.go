package zgen

// NoCap returns the slice provided except all excess cap is removed.
func NoCap[T any](s []T) []T {
	return s[:len(s):len(s)]
}
