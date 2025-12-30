package stack

type Stack[T any] struct {
	data []T
}

// New returns an empty stack.
// Optional, you could just do `var s stack.Stack[T]{}`.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Push adds many element with the rightmost element to the top of the stack.
func (s *Stack[T]) PushMany(v ...T) {
	s.data = append(s.data, v...)
}

// Pop removes and returns the top element of the stack.
//
// If the stack is empty, Pop returns the zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	size := len(s.data)
	if size == 0 {
		var zero T
		return zero, false
	}

	// Get element
	v := s.data[size-1]

	// Clear references
	var zero T
	s.data[size-1] = zero

	// Shrink stack len
	s.data = s.data[:size-1]

	return v, true
}

// PopN removes N elements and returns them in stack order, where out[0] is the top.
//
// If the stack is empty, Pop returns nil and false.
func (s *Stack[T]) PopN(n int) ([]T, bool) {
	size := len(s.data)
	if size == 0 || n > size {
		return nil, false
	}

	start := size - n

	// Copy out values so rather than a view it is actually own by them
	out := make([]T, n)
	for i := 0; i < n; i++ {
		out[i] = s.data[size-1-i]
	}

	// Clear references
	var zero T
	for i := start; i < size; i++ {
		s.data[i] = zero
	}

	// Shrink stack len
	s.data = s.data[:start]

	return out, true
}

// Peek returns the top element without removing it.
//
// If the stack is empty, Peek returns the zero value of T and false.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// IsEmpty returns boolean whether the stack has no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	// Avoid holding references
	for i := range s.data {
		var zero T
		s.data[i] = zero
	}
	s.data = s.data[:0]
}
