package stack

import "testing"

func TestStack_PushPop(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Fatalf("expected stack to be empty")
	}

	s.Push(10)
	s.Push(20)

	if size := s.Size(); size != 2 {
		t.Fatalf("expected size 2, got %d", size)
	}

	v, ok := s.Pop()
	if !ok || v != 20 {
		t.Fatalf("expected pop to return (20, true); got %v, %v", v, ok)
	}

	v, ok = s.Pop()
	if !ok || v != 10 {
		t.Fatalf("expected pop to return (10, true); got %v, %v", v, ok)
	}

	if _, ok := s.Pop(); ok {
		t.Fatalf("expected pop on empty stack to return false")
	}
}

func TestStack_PushManyPopN(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Fatalf("expected stack to be empty")
	}

	s.PushMany([]int{10, 20}...)

	if size := s.Size(); size != 2 {
		t.Fatalf("expected size 2, got %d", size)
	}

	v, ok := s.PopN(2)
	expect := []int{10, 20}

	if !ok {
		t.Fatalf("expected PopN to return ok=true")
	}

	if len(v) != len(expect) {
		t.Fatalf("expected %d elements, got %d", len(expect), len(v))
	}

	for i := range expect {
		if v[i] != expect[i] {
			t.Fatalf("expected %v, got %v", expect, v)
		}
	}

	if _, ok := s.Pop(); ok {
		t.Fatalf("expected pop on empty stack to return false")
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack[string]{}

	if _, ok := s.Peek(); ok {
		t.Fatalf("expected peek on empty stack to fail")
	}

	s.Push("a")
	s.Push("b")

	v, ok := s.Peek()
	if !ok || v != "b" {
		t.Fatalf("expected peek to return 'b', true; got %q, %v", v, ok)
	}

	// Peek must not remove the element
	if size := s.Size(); size != 2 {
		t.Fatalf("expected size 2 after peek, got %d", size)
	}
}

func TestStack_ZeroValue(t *testing.T) {
	var s Stack[int]

	if !s.IsEmpty() {
		t.Fatalf("zero value stack should be empty")
	}

	s.Push(42)

	v, ok := s.Pop()
	if !ok || v != 42 {
		t.Fatalf("expected pop to return 42, true; got %v, %v", v, ok)
	}
}

func TestStack_GenericTypes(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}

	s := New[user]()

	u := user{ID: 1, Name: "Alice"}
	s.Push(u)

	got, ok := s.Pop()
	if !ok || got != u {
		t.Fatalf("expected %v, got %v", u, got)
	}
}
