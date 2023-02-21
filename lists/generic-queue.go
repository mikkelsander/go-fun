package lists

import (
	"errors"
	"fmt"
)

type GenericQueue[T any] struct {
	slice []T
}

func (q *GenericQueue[T]) Enqueue(item T) {
	q.slice = append(q.slice, item)
}

func (q *GenericQueue[T]) Dequeue() (T, error) {
	if len(q.slice) > 0 {
		item := q.slice[0]
		q.slice = q.slice[1:]
		return item, nil
	} else {
		return *new(T), errors.New("queue is empty")
	}
}

func (q *GenericQueue[T]) Peek() T {
	return q.slice[0]
}

func (q *GenericQueue[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

func (q *GenericQueue[T]) Print() {
	fmt.Printf("%v", q.slice)
}

func (q *GenericQueue[T]) Push(item T) {
	q.slice = append(q.slice, item)
}

func (q *GenericQueue[T]) Pop() (T, error) {
	if len(q.slice) > 0 {
		var item = q.slice[len(q.slice)-1]
		q.slice = q.slice[:len(q.slice)-1]
		return item, nil
	} else {
		return *new(T), errors.New("queue is empty")
	}
}
