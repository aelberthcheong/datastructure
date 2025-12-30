package queue

type Queue[T any] struct {
	data []T
	head int // index of the front element
	size int // number of elements currently in the queue
}

// New creates a new queue with the given initial capacity.
// Capacity is automatically increased as needed.
func New[T any](capacity int) *Queue[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &Queue[T]{
		data: make([]T, capacity),
	}
}

// Len returns the number of elements currently in the queue.
func (q *Queue[T]) Len() int {
	return q.size
}

// Cap returns the current capacity of the queue.
func (q *Queue[T]) Cap() int {
	return len(q.data)
}

// IsEmpty reports whether the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	if q.size == len(q.data) {
		q.grow()
	}

	idx := (q.head + q.size) % len(q.data)
	q.data[idx] = v
	q.size++
}

// Dequeue removes and returns the front element of the queue.
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	v := q.data[q.head]

	// Clear slot to avoid retaining references
	var zero T
	q.data[q.head] = zero

	q.head = (q.head + 1) % len(q.data)
	q.size--

	// Optional: normalize head when empty
	if q.size == 0 {
		q.head = 0
	}

	return v, true
}

// Peek returns the front element without removing it.
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Peek() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.data[q.head], true
}

// Clear removes all elements from the queue while retaining capacity.
func (q *Queue[T]) Clear() {
	var zero T
	for i := range q.data {
		q.data[i] = zero
	}
	q.head = 0
	q.size = 0
}

// grow doubles the capacity of the queue while preserving element order.
func (q *Queue[T]) grow() {
	oldBuf := q.data
	oldCap := len(oldBuf)
	newCap := oldCap * 2

	if newCap == 0 {
		newCap = 1
	}

	newBuf := make([]T, newCap)

	// Copy elements in logical order
	for i := 0; i < q.size; i++ {
		newBuf[i] = oldBuf[(q.head+i)%oldCap]
	}

	q.data = newBuf
	q.head = 0
}
