package unordered_map

type UnorderedMap[T1 comparable, T2 any] map[T1]T2

func New[T1 comparable, T2 any]() *UnorderedMap[T1, T2] {
	return &UnorderedMap[T1, T2]{}
}

func (m *UnorderedMap[T1, T2]) Insert(key T1, value T2) {
	(*m)[key] = value
}

func (m *UnorderedMap[T1, T2]) Erase(key T1) {
	delete(*m, key)
}

func (m *UnorderedMap[T1, T2]) Find(key T1) (T2, bool) {
	value, ok := (*m)[key]
	return value, ok
}

func (m *UnorderedMap[T1, T2]) Contain(key T1) bool {
	_, ok := (*m)[key]
	return ok
}

func (m *UnorderedMap[T1, T2]) Empty() bool {
	return len(*m) == 0
}

func (m *UnorderedMap[T1, T2]) Size() int {
	return len(*m)
}
