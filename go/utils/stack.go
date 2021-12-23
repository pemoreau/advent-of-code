package utils

import "errors"

type Stack []interface{}

func BuildStack() Stack {
	return make([]interface{}, 0)
}

func (s *Stack) Push(c interface{}) {
	*s = append(*s, c)
}

func (s *Stack) Pop() (interface{}, error) {
	l := len(*s)
	if l == 0 {
		return 0, errors.New("stack is empty")
	}
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
