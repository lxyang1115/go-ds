package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) Front() T {
	return (*q)[0]
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Clear() {
	*q = (*q)[:0]
}
