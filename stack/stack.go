// The stack is not safe for concurrent use.
package stack

type Stack[T any] struct {
	data []T
}

// New returns an empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top element of the stack.
//
// If the stack is empty, Pop returns the zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	n := len(s.data)
	if n == 0 {
		var zero T
		return zero, false
	}

	// Get element
	v := s.data[n-1]

	// Prevent memory leaks for reference types
	var zero T
	s.data[n-1] = zero

	// Shrink slice
	s.data = s.data[:n-1]

	return v, true
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

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
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
