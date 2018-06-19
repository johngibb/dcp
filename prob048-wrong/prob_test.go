package prob

import (
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

func TestFromPreorder(t *testing.T) {
	got := TreeFromPreorder([]string{"a", "b", "d", "e", "c", "f", "g"})

	var (
		gotWalk  = Preorder(got)
		wantWalk = Preorder(exampleTree)
	)
	if !reflect.DeepEqual(gotWalk, wantWalk) {
		t.Errorf("got %v, want %v", gotWalk, wantWalk)
	}
}

func TestFromInorder(t *testing.T) {
	got := TreeFromInorder([]string{"d", "b", "e", "a", "f", "c", "g"})

	var (
		gotWalk  = Preorder(got)
		wantWalk = Preorder(exampleTree)
	)
	if !reflect.DeepEqual(gotWalk, wantWalk) {
		t.Errorf("got %v, want %v", gotWalk, wantWalk)
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
