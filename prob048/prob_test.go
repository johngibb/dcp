package prob

import (
	"fmt"
	"reflect"
	"testing"
)

var exampleTree = &Node{
	Val: "a",
	Left: &Node{
		Val:   "b",
		Left:  &Node{Val: "d"},
		Right: &Node{Val: "e"},
	},
	Right: &Node{
		Val:   "c",
		Left:  &Node{Val: "f"},
		Right: &Node{Val: "g"},
	},
}

var tests []*Node = []*Node{
	exampleTree,
	&Node{
		Val:  "a",
		Left: &Node{Val: "b"},
	},
	&Node{
		Val:   "a",
		Right: &Node{Val: "b"},
	},
}

func Test(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("tree %d", i), func(t *testing.T) {
			var (
				pre  = Preorder(tt)
				in   = Inorder(tt)
				tree = BuildTree(pre, in)
				want = print(tt)
				got  = print(tree)
			)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func Test2(t *testing.T) {
	tests := []struct {
		pre, in []string
	}{
		{
			[]string{"a", "b", "d", "e", "c", "f"},
			[]string{"d", "b", "e", "a", "f", "c"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := BuildTree(tt.pre, tt.in)
			if got := Preorder(tree); !reflect.DeepEqual(got, tt.pre) {
				t.Errorf("got %v, want %v", got, tt.pre)
			}
			if got := Inorder(tree); !reflect.DeepEqual(got, tt.in) {
				t.Errorf("got %v, want %v", got, tt.in)
			}
		})
	}
}

func TestPreorder(t *testing.T) {
	want := []string{"a", "b", "d", "e", "c", "f", "g"}
	got := Preorder(exampleTree)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestInorder(t *testing.T) {
	want := []string{"d", "b", "e", "a", "f", "c", "g"}
	got := Inorder(exampleTree)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func print(n *Node) string {
	var rec func(n *Node) string
	rec = func(n *Node) string {
		if n == nil {
			return "_"
		}
		if n.leaf() {
			return n.Val
		}
		var s string
		s += "(" + rec(n.Left) + ")"
		s += n.Val
		s += "(" + rec(n.Right) + ")"
		return s
	}
	return rec(n)
}
