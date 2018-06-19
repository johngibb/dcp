package prob

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

func (n *Node) leaf() bool {
	return n.Left == nil && n.Right == nil
}

func BuildTree(preorder, inorder []string) *Node {
	var i int
	indexOf := func(vals []string, v string) int {
		for i, vv := range vals {
			if v == vv {
				return i
			}
		}
		panic("not found")
	}
	var rec func(start, end int) *Node
	rec = func(start, end int) *Node {
		if start > end {
			return nil
		}
		v := preorder[i]
		i++
		n := &Node{Val: v}
		if start <= end {
			idx := indexOf(inorder, v)
			n.Left = rec(start, idx-1)
			n.Right = rec(idx+1, end)
		}
		return n
	}
	return rec(0, len(preorder)-1)
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
