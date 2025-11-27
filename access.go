package zgen

func last2[T2 any](_ any, v2 T2) T2 { return v2 }

func BoolLast2(_ any, b bool) bool { return last2(nil, b) }

func ErrLast2(_ any, err error) error { return last2(nil, err) }
