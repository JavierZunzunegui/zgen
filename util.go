package zgen

// TODO - document.
func PtrTo[T any](t T) *T {
	return &t
}
