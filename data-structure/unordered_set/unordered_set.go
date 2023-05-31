package unordered_set

type UnorderedSet[T comparable] map[T]struct{}

func New[T comparable](es ...T) *UnorderedSet[T] {
	s := &UnorderedSet[T]{}
	for _, e := range es {
		s.Insert(e)
	}
	return s
}

func (s *UnorderedSet[T]) Insert(key T) {
	(*s)[key] = struct{}{}
}

func (s *UnorderedSet[T]) Erase(key T) {
	delete(*s, key)
}

func (s *UnorderedSet[T]) Contain(key T) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *UnorderedSet[T]) Empty() bool {
	return len(*s) == 0
}

func (s *UnorderedSet[T]) Size() int {
	return len(*s)
}

func (s *UnorderedSet[T]) Clear() {
	*s = map[T]struct{}{}
}

func (s *UnorderedSet[T]) ToSlice() []T {
	keys := make([]T, 0, len(*s))
	for key := range *s {
		keys = append(keys, key)
	}
	return keys
}
