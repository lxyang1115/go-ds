package main

import (
	"fmt"
	"go-ds/data-structure/unordered_set"
)

func main() {
	s := unordered_set.New[int](1, 2, 3)
	fmt.Println(s.Contain(1))
	fmt.Println(s.Contain(4))
	s.Erase(1)
	fmt.Println(len(s.ToSlice()))
	fmt.Println(s.Contain(1))
	fmt.Println(s.Size())
	s.Clear()
	fmt.Println(s.Empty())
}
