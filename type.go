package zgen

// PtrTo is a lazy way to do `var x T; &x`, but do so in a single statement.
func PtrTo[T any](t T) *T {
	return &t
}

// IsType is a lazy way to check if the value in an interface can be cast to another type.
// Note if [T] is an interface the check is for 'can cast', if it is a concrete type the check
// is for type equality.
// See also the very similar [Cast].
//
// IsType[T](x) <=> {_, ok := x.(T); return ok}
func IsType[T any](i any) bool {
	_, ok := i.(T)
	return ok
}

// Cast is similar to [IsType] except it also returns the type it was casted to.
// Note when returning false, the casted value is the T zero value.
//
// Cast[T](x) <=> {t, ok := x.(T); return t, ok}
func Cast[T any](i any) (T, bool) {
	t, ok := i.(T)
	return t, ok
}
