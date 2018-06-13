package prob003

import (
	"strconv"
	"strings"
)

// Given the root to a binary tree, implement serialize(root), which serializes
// the tree into a string, and deserialize(s), which deserializes the string
// back into the tree.

type Node struct {
	Val         int
	Left, Right *Node
}

func SerializeArray(root *Node) string {
	var sb strings.Builder
	var rec func(n *Node)
	rec = func(n *Node) {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}
		if n == nil {
			sb.WriteString("_")
			return
		}
		sb.WriteString(strconv.Itoa(n.Val))
		rec(n.Left)
		rec(n.Right)
	}
	rec(root)
	return sb.String()
}

func DeserializeArray(s string) (n *Node, err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			err = err2.(error)
		}
	}()
	vals := strings.Split(s, ",")
	i := -1
	var rec func() *Node
	rec = func() *Node {
		i++
		if vals[i] == "_" {
			return nil
		}
		val, err := strconv.Atoi(vals[i])
		if err != nil {
			panic(err)
		}
		return &Node{
			Val:   val,
			Left:  rec(),
			Right: rec(),
		}
	}
	return rec(), nil
}
