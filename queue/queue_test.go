package queue

import "testing"

func TestQueue_New(t *testing.T) {
	q := New[int](1)

	if q.Len() != 0 {
		t.Fatalf("expected empty queue, got len=%d", q.Len())
	}

	if q.Cap() < 1 {
		t.Fatalf("expected capacity >= 1, got %d", q.Cap())
	}
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	var q Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Fatalf("expected len=2, got %d", q.Len())
	}

	v, ok := q.Dequeue()
	if !ok || v != 1 {
		t.Fatalf("expected (1, true), got (%v, %v)", v, ok)
	}

	v, ok = q.Dequeue()
	if !ok || v != 2 {
		t.Fatalf("expected (2, true), got (%v, %v)", v, ok)
	}

	if _, ok := q.Dequeue(); ok {
		t.Fatalf("expected dequeue on empty queue to fail")
	}
}

func TestQueue_Peek(t *testing.T) {
	var q Queue[int]

	if _, ok := q.Peek(); ok {
		t.Fatalf("expected peek on empty queue to fail")
	}

	q.Enqueue(42)

	v, ok := q.Peek()
	if !ok || v != 42 {
		t.Fatalf("expected peek to return (42, true), got (%v, %v)", v, ok)
	}

	if q.Len() != 1 {
		t.Fatalf("peek should not remove element")
	}
}

func TestQueue_WrapAround(t *testing.T) {
	var q Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Force head movement
	v, _ := q.Dequeue()
	if v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	q.Enqueue(4)

	expected := []int{2, 3, 4}
	for _, exp := range expected {
		v, ok := q.Dequeue()
		if !ok || v != exp {
			t.Fatalf("expected %d, got (%v, %v)", exp, v, ok)
		}
	}
}

func TestQueue_Grow(t *testing.T) {
	var q Queue[int]
	initialCap := q.Cap()

	// Force growth
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	if q.Cap() <= initialCap {
		t.Fatalf("expected capacity to grow beyond %d, got %d", initialCap, q.Cap())
	}

	for i := 0; i < 10; i++ {
		v, ok := q.Dequeue()
		if !ok || v != i {
			t.Fatalf("expected %d, got (%v, %v)", i, v, ok)
		}
	}
}

func TestQueue_Clear(t *testing.T) {
	var q Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)

	q.Clear()

	if q.Len() != 0 {
		t.Fatalf("expected empty queue after Clear, got len=%d", q.Len())
	}

	if _, ok := q.Dequeue(); ok {
		t.Fatalf("expected dequeue after Clear to fail")
	}

	// Ensure queue is reusable
	q.Enqueue(3)
	v, ok := q.Dequeue()
	if !ok || v != 3 {
		t.Fatalf("expected queue to be reusable after Clear")
	}
}
