package prob003

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	trees := []*Node{
		&Node{
			Val:   1,
			Left:  &Node{Val: 2},
			Right: &Node{Val: 3},
		},
		&Node{
			Val: 2341,
			Left: &Node{
				Val:  210,
				Left: &Node{Val: 937},
			},
			Right: &Node{Val: 3},
		},
	}
	strategies := []struct {
		Name        string
		Serialize   func(root *Node) string
		Deserialize func(s string) (*Node, error)
	}{
		{"Array", SerializeArray, DeserializeArray},
	}

	for _, s := range strategies {
		t.Run(s.Name, func(t *testing.T) {
			for i, tt := range trees {
				t.Run(fmt.Sprintf("Tree %d", i), func(t *testing.T) {
					serialized := s.Serialize(tt)
					deserialized, err := s.Deserialize(serialized)

					if err != nil {
						t.Errorf("error deserializing: %v", err)
					}

					if !reflect.DeepEqual(tt, deserialized) {
						t.Errorf("got %v, want %v", deserialized, tt)
					}
				})
			}
		})
	}
}
