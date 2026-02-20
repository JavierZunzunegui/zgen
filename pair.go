package zgen

type Pair[A, B any] struct {
	A A
	B B
}

func (p Pair[A, B]) First() A     { return p.A }
func (p Pair[A, B]) Second() B    { return p.B }
func (p Pair[A, B]) Both() (A, B) { return p.A, p.B }

func NewPair[A, B any](a A, b B) Pair[A, B] { return Pair[A, B]{A: a, B: b} }
