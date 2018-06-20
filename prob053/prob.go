package prob

// This problem was asked by Apple.
//
// Implement a queue using two stacks. Recall that a queue is a FIFO (first-in,
// first-out) data structure with the following methods: enqueue, which inserts
// an element into the queue, and dequeue, which removes it.

type Queue struct {
	a, b Stack
}

func (q *Queue) Enqueue(v string) {
	q.a.Push(v)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.a.Len() == 0 {
		return "", false
	}
	// Reverse stack a on to stack b.
	v, ok := q.a.Pop()
	for ok {
		q.b.Push(v)
		v, ok = q.a.Pop()
	}

	// Pop the return value from stack b.
	result, ok := q.b.Pop()
	if !ok {
		panic("no result; race condition?")
	}

	// Undo the reverse.
	v, ok = q.b.Pop()
	for ok {
		q.a.Push(v)
		v, ok = q.b.Pop()
	}

	return result, true
}

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return v, true
}

func (s Stack) Len() int {
	return len(s)
}
