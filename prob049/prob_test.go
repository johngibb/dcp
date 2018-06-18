package prob

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{34, -50, 42, 14, -5, 86}, 137},
		{[]int{-5, -1, -8, -0}, 0},
	}
	strategies := []struct {
		name string
		fn   func([]int) int
	}{
		{"MaxContiguousSum", MaxContiguousSum},
		{"MaxContiguousSumKadane", MaxContiguousSumKadane},
	}
	for _, s := range strategies {
		t.Run(s.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
					got := s.fn(tt.in)
					if got != tt.want {
						t.Errorf("got %d, want %d", got, tt.want)
					}
				})
			}
		})
	}
}
