package prob004

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{3, 4, -1, 1}, 2},
		{[]int{1, 2, 3, 4}, 5},
		{[]int{1, 2, 0}, 3},
		{[]int{1, 2, 2, 0}, 3},
	}

	strategies := []struct {
		name string
		fn   func([]int) int
	}{
		{"Naive", FindMissingPositiveIntegerNaive},
		{"Efficient", FindMissingPositiveIntegerEfficient},
		{"Alternate", FindMissingPositiveIntegerAlternate},
	}

	for _, strategy := range strategies {
		t.Run(strategy.name, func(t *testing.T) {
			for i, tt := range tests {
				t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
					got := strategy.fn(clone(tt.in))
					if got != tt.want {
						t.Errorf("%v: got %d, want %d", tt.in, got, tt.want)
					}
				})
			}
		})
	}
}

func clone(vals []int) []int {
	cp := make([]int, len(vals))
	for i, v := range vals {
		cp[i] = v
	}
	return cp
}
