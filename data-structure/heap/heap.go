package heap

type Interface[T any] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(v T)
	Pop() T
}
