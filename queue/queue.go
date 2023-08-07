package queue

type constraint interface {
	int | string
}

type Queue[T constraint] struct {
	vals []T
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.vals) == 0 {
		var zero T
		return zero, false
	}
	bottom := q.vals[0]
	q.vals = q.vals[1:]
	return bottom, true
}

func (q *Queue[T]) Enqueue(val ...T) {
	q.vals = append(q.vals, val...)
}
