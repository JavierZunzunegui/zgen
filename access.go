package zgen

func last2[T2 any](_ any, v2 T2) T2 { return v2 }

// BoolLast2 discards its first argument and returns the bool second argument.
// It is intended for wrapping two-return-value calls where only the bool matters:
//
//	ok := zgen.BoolLast2(f())
func BoolLast2(_ any, b bool) bool { return last2(nil, b) }

// ErrLast2 discards its first argument and returns the error second argument.
// It is intended for wrapping two-return-value calls where only the error matters:
//
//	err := zgen.ErrLast2(f())
func ErrLast2(_ any, err error) error { return last2(nil, err) }
