package prob

import (
	"testing"
)

func Test(t *testing.T) {
	tree := &Node{
		Val:  "a",
		Left: &Node{Val: "c"},
		Right: &Node{
			Val:  "b",
			Left: &Node{Val: "b"},
			Right: &Node{
				Val:   "b",
				Right: &Node{Val: "b"},
			},
		},
	}
	got := CountUnivalSubtrees(tree)
	if got != 5 {
		t.Errorf("got: %d, want 5", got)
	}
}
