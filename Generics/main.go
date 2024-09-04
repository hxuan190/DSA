package main

import "fmt"

// Định nghĩa kiểu dữ liệu Stack với type parameter T
type Stack[T any] struct {
	elements []T
}

// Phương thức Push thêm phần tử vào Stack
func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() T {
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println(intStack.Pop()) // Output: 20

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	fmt.Println(stringStack.Pop()) // Output: World
}
