package pkg

type Queue[T any] struct {
	contents []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		contents: make([]T, 0, 5),
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}

func (q *Queue[T]) Enqueue(item T) {
	q.contents = append(q.contents, item)
}

func (q *Queue[T]) Dequeue() T {
	firstElement := q.contents[0]

	q.contents = q.contents[1:]
	return firstElement
}
