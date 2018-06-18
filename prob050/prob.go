package prob

import (
	"fmt"
	"strconv"
)

// This problem was asked by Microsoft.
//
// Suppose an arithmetic expression is given as a binary tree. Each leaf is an
// integer and each internal node is one of '+', '−', '∗', or '/'.
//
// Given the root to such a tree, write a function to evaluate it.
//
// For example, given the following tree:
//
//     *
//    / \
//   +    +
//  / \  / \
// 3  2  4  5
// You should return 45, as it is (3 + 2) * (4 + 5).

type Node struct {
	Val         int
	Left, Right *Node
}

var ops = []func(int, int) int{
	func(a, b int) int { return a + b },
	func(a, b int) int { return a - b },
	func(a, b int) int { return a * b },
	func(a, b int) int { return a / b },
}

const (
	OpAdd      = 0
	OpSubtract = 1
	OpMultiply = 2
	OpDivide   = 3
)

func (n *Node) leaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) String() string {
	if n.leaf() {
		return strconv.Itoa(n.Val)
	} else {
		switch n.Val {
		case OpAdd:
			return "+"
		case OpSubtract:
			return "-"
		case OpMultiply:
			return "*"
		case OpDivide:
			return "/"
		default:
			panic(fmt.Sprintf("invalid internal node: %d", n.Val))
		}
	}
}

func Eval(n *Node) int {
	var rec func(n *Node) int
	rec = func(n *Node) int {
		if n.leaf() {
			return n.Val
		}
		op := ops[n.Val]
		return op(rec(n.Left), rec(n.Right))
	}
	return rec(n)
}
