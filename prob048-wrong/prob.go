package prob

import (
	"math"
)

// NOTE: This is wrong. I misread the question, thinking it was asking for a way
// to reconstruct a _full_ binary tree from just the pre-order, and just the in-
// order walks. For non-full trees, this isn't possible.

// Given pre-order and in-order traversals of a binary tree, write a function to
// reconstruct the tree.

// For example, given the following preorder traversal:

// [a, b, d, e, c, f, g]

// And the following inorder traversal:

// [d, b, e, a, f, c, g]

// You should return the following tree:

//     a
//    / \
//   b   c
//  / \ / \
// d  e f  g

type Node struct {
	Val         string
	Left, Right *Node
}

func TreeFromPreorder(vals []string) *Node {
	i := 0
	var rec func() *Node
	rec = func() *Node {
		if i >= len(vals) {
			return nil
		}
		val := vals[i]
		i++
		return &Node{
			Val:   val,
			Left:  rec(),
			Right: rec(),
		}
	}
	return rec()
}

func TreeFromInorder(vals []string) *Node {
	// Always going to be an odd number of nodes.
	depth := int(math.Sqrt(float64(len(vals))))

	var rec func(d, i int) *Node
	rec = func(d, i int) *Node {
		if i < 0 || i >= len(vals) {
			return nil
		}
		n := &Node{
			Val: vals[i],
		}
		if d > 0 {
			n.Left = rec(d-1, i-d)
			n.Right = rec(d-1, i+d)
		}
		return n
	}
	return rec(depth, len(vals)/2)
}

func Preorder(n *Node) []string {
	var result []string
	var rec func(n *Node)
	rec = func(n *Node) {
		if n == nil {
			return
		}
		result = append(result, n.Val)
		rec(n.Left)
		rec(n.Right)
	}
	rec(n)
	return result
}

func Inorder(n *Node) []string {
	var result []string
	var rec func(n *Node)
	rec = func(n *Node) {
		if n == nil {
			return
		}
		rec(n.Left)
		result = append(result, n.Val)
		rec(n.Right)
	}
	rec(n)
	return result
}
