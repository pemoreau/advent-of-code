package stack

import "errors"

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func (s *Stack[T]) Push(c T) {
	*s = append(*s, c)
}

func (s *Stack[T]) Pop() (T, error) {
	l := len(*s)
	if l == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top, nil
}

func (s *Stack[T]) PushN(c []T) {
	*s = append(*s, c...)
}

func (s *Stack[T]) PopN(n int) ([]T, error) {
	l := len(*s)
	if l < n {
		var zero []T
		return zero, errors.New("stack is empty")
	}
	top := (*s)[l-n:]
	*s = (*s)[:l-n]
	return top, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
