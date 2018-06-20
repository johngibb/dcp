package prob

// Good morning. Here's your coding interview problem for today.
// This problem was asked by Google.
//
// A unival tree (which stands for "universal value") is a tree where all nodes
// have the same value.
//
// Given the root to a binary tree, count the number of unival subtrees.

type Node struct {
	Val         string
	Left, Right *Node
}

func CountUnivalSubtrees(root *Node) int {
	var rec func(*Node) (int, bool)
	rec = func(n *Node) (int, bool) {
		var (
			unival = true
			count  = 0
		)
		for _, child := range []*Node{n.Left, n.Right} {
			if child == nil {
				continue
			}
			// A node is unival if all of its children are unival, and have the
			// same value as itself.
			childCount, childUnival := rec(child)
			if child.Val != n.Val || !childUnival {
				unival = false
			}
			count += childCount
		}
		if unival {
			count += 1
		}
		return count, unival
	}
	result, _ := rec(root)
	return result
}
