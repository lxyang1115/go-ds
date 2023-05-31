package pair

type Pair struct {
	First  any
	Second any
}

func MakePair(first any, second any) *Pair {
	return &Pair{First: first, Second: second}
}

func (p *Pair) GetFirst() any {
	return p.First
}

func (p *Pair) GetSecond() any {
	return p.Second
}

func (p *Pair) SetFirst(first any) {
	p.First = first
}

func (p *Pair) SetSecond(second any) {
	p.Second = second
}

func (p *Pair) Equal(other *Pair) bool {
	return p.First == other.First && p.Second == other.Second
}
