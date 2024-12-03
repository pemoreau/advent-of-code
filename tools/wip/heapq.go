package wip

import "fmt"

type Heap[T any] struct { // accept Any type into the queue
	// compare two elements of type T and return true if first is smaller
	comparator Comparator[T]
	// number of elements in the priority queue
	size int
	// memory buffer of fixed capacity for storing the binary tree of items
	data []T
}

// return true if a < b
type Comparator[T any] func(a *T, b *T) bool

func NewHeap[T any](capacity int, comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
		size:       0,
		data:       make([]T, capacity+1, capacity+1),
	}
}

func parentIdx(pos int) int {
	return (pos - 1) / 2
}

func leftIdx(pos int) int {
	return 2*pos + 1
}

func rightIdx(pos int) int {
	return 2*pos + 2
}

func (q *Heap[T]) isLeaf(pos int) bool {
	return leftIdx(pos) > q.size
}

func (q *Heap[T]) swap(a int, b int) {
	q.data[a], q.data[b] = q.data[b], q.data[a]
}

func (q *Heap[T]) Peek() (res T, err error) {
	if q.size < 1 {
		return res, fmt.Errorf("peeking into an empty queue")
	}

	res = q.data[0]
	return res, nil
}

func (q *Heap[T]) Push(item T) error {
	if q.size >= len(q.data) {
		return fmt.Errorf("pushing into a full container")
	}

	cur := q.size
	q.size++

	q.data[cur] = item
	for q.comparator(&q.data[cur], &q.data[parentIdx(cur)]) {
		q.swap(cur, parentIdx(cur))
		cur = parentIdx(cur)
	}

	return nil
}

func (q *Heap[T]) Pop() (res T, err error) {
	if q.size < 1 {
		return res, fmt.Errorf("popping from an empty queue")
	}

	if q.size == 1 {
		q.size--
		return q.data[0], nil
	}

	res = q.data[0]
	q.data[0] = q.data[q.size-1]
	q.size--
	q.percolate(0)

	return res, nil
}

func (q *Heap[T]) percolate(pos int) {
	if q.isLeaf(pos) {
		return
	}

	var cur *T = &q.data[pos]
	var left *T = &q.data[leftIdx(pos)]
	var right *T
	if rightIdx(pos) <= q.size {
		right = &q.data[rightIdx(pos)]
	}

	if q.comparator(left, cur) || q.comparator(right, cur) {
		if q.comparator(left, right) {
			q.swap(pos, leftIdx(pos))
			q.percolate(leftIdx(pos))
		} else {
			q.swap(pos, rightIdx(pos))
			q.percolate(rightIdx(pos))
		}
	}
}
