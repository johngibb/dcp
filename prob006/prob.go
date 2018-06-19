package prob006

import "strconv"

// An XOR linked list is a more memory efficient doubly linked list. Instead of
// each node holding next and prev fields, it holds a field named both, which is
// a XOR of the next node and the previous node. Implement a XOR linked list; it
// has an add(element) which adds the element to the end, and a get(index) which
// returns the node at index.
//
// If using a language that has no pointers (such as Python), assume you have
// access to get_pointer and dereference_pointer functions that converts between
// nodes and memory addresses.

type Node struct {
	Val        int
	Next, Prev *Node
}

func (n *Node) SetLinks(next, prev *Node) {
	n.Next = next
	n.Prev = prev
}

func (n *Node) GetNext(prev *Node) *Node {
	return n.Next
}

func (n *Node) Add(v int) {
	if n.Next != nil {
		n.Next.Add(v)
		return
	}
	next := &Node{
		Val: v,
	}
	next.SetLinks(nil, n)
	n.SetLinks(next, n.Prev)
}

func (n *Node) Get(i int) int {
	for i > 0 {
		n = n.Next
		i--
	}
	return n.Val
}

func (n *Node) String() string {
	s := strconv.Itoa(n.Val)
	if n.Next != nil {
		s += " -> " + n.Next.String()
	}
	return s
}
