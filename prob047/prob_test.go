package prob

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		profit int
	}{
		{[]int{1, 5, 0, 10}, 10},
		{[]int{9, 11, 8, 5, 7, 10}, 5},
		{[]int{2, 7, 1, 8, 2, 8, 4, 5, 9, 0, 4, 5}, 8},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.profit {
				t.Errorf("got %d, want %d", got, tt.profit)
			}
		})
	}
}
