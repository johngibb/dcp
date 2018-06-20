package prob

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var q Queue
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	for i, want := range []string{"a", "b", "c"} {
		v, ok := q.Dequeue()
		if !ok {
			t.Error("couldn't dequeue")
		}
		if v != want {
			t.Errorf("[%d] q.Dequeue(): got %q, want %q", i, v, want)
		}
	}

	if _, ok := q.Dequeue(); ok {
		t.Error("q.Dequeue() should return ok=false for empty queue")
	}
}

func TestStack(t *testing.T) {
	var stack Stack

	if stack.Len() != 0 {
		t.Error("stack.Len() should be zero for a new struct")
	}

	stack.Push("test")

	if len := stack.Len(); len != 1 {
		t.Errorf("stack.Len(): got %d, want 1", len)
	}

	v, ok := stack.Pop()
	if want := "test"; v != want {
		t.Errorf("stack.Pop(): got %q, want %q", v, want)
	}
	if !ok {
		t.Error("stack.Pop(): returned ok=false when not empty")
	}

	v, ok = stack.Pop()
	if v != "" {
		t.Errorf("stack.Pop(): got %q, want empty", v)
	}
	if ok {
		t.Error("stack.Pop() returned ok=false on non empty-stack")
	}
}
