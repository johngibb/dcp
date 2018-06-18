package prob

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		tree *Node
		want int
	}{
		{
			tree: &Node{
				Val: OpMultiply,
				Left: &Node{
					Val:   OpAdd,
					Left:  &Node{Val: 3},
					Right: &Node{Val: 2},
				},
				Right: &Node{
					Val:   OpAdd,
					Left:  &Node{Val: 4},
					Right: &Node{Val: 5},
				},
			},
			want: 45,
		},
		{
			tree: &Node{
				Val: OpDivide,
				Left: &Node{
					Val:   OpMultiply,
					Left:  &Node{Val: 10},
					Right: &Node{Val: 10},
				},
				Right: &Node{
					Val:   OpAdd,
					Left:  &Node{Val: 9},
					Right: &Node{Val: 1},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(printTree(tt.tree), func(t *testing.T) {
			got := Eval(tt.tree)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func printTree(n *Node) string {
	var rec func(n *Node) string
	rec = func(n *Node) string {
		if n.leaf() {
			return n.String()
		}
		var s string
		if n.Left != nil {
			s += rec(n.Left)
		}
		s += n.String()
		if n.Right != nil {
			s += rec(n.Right)
		}
		return "(" + s + ")"
	}
	return rec(n)
}
