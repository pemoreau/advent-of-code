package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) AddAll(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Values() []T {
	res := make([]T, 0, len(s))
	for value := range s {
		res = append(res, value)
	}
	return res
}

func (s Set[T]) Element() T {
	for elem := range s {
		return elem
	}
	panic("Set is empty")
}

func (s Set[T]) Pop() T {
	for elem := range s {
		delete(s, elem)
		return elem
	}
	panic("Set is empty")
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Clear() {
	clear(s)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Equal(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	res := make(Set[T], max(s.Len(), other.Len()))
	for elem := range s {
		res.Add(elem)
	}
	for elem := range other {
		res.Add(elem)
	}
	return res
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	res := NewSet[T]()
	if s.Len() < other.Len() {
		for elem := range s {
			if other.Contains(elem) {
				res.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if s.Contains(elem) {
				res.Add(elem)
			}
		}
	}
	return res
}

func (s Set[T]) MutableIntersect(other Set[T]) {
	for elem := range s {
		if !other.Contains(elem) {
			delete(s, elem)
		}
	}
}
