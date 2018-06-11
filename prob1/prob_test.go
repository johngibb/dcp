package prob1

import "testing"

// Given an array of numbers, return whether any two sums to K.
//
// For example, given [10, 15, 3, 7] and K of 17, return true since 10 + 7 is
// 17.

func TestProb1(t *testing.T) {
	tests := []struct {
		vals []int
		sum  int
		want bool
	}{
		{[]int{10, 15, 3, 7}, 17, true},
		{[]int{1, 2, 3}, 5, true},
		{[]int{1, 2, 3}, 6, false},
		{[]int{1, 2}, 4, false},
		{[]int{2, 2}, 4, true},
	}
	strategies := []struct {
		fn   func([]int, int) bool
		name string
	}{
		{ContainsSumPairNaive, "Naive"},
		{ContainsSumPairMap, "Map"},
	}
	for _, tt := range tests {
		for _, s := range strategies {
			got := s.fn(tt.vals, tt.sum)
			if got != tt.want {
				t.Errorf("%s: vals: %v, sum: %d: got %v, want %v", s.name, tt.vals, tt.sum, got, tt.want)
			}
		}
	}
}
