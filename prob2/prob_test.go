package prob2

import (
	"reflect"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{120, 60, 40, 30, 24},
		},
		{
			[]int{3, 2, 1},
			[]int{2, 3, 6},
		},
	}
	strategies := []struct {
		fn   func(input []int) []int
		name string
	}{
		{ComputeNaive, "Naive"},
		{ComputeEfficient, "Efficient"},
		{ComputeEfficientOneArray, "EfficientOneArray"},
	}
	for _, s := range strategies {
		for _, tt := range tests {
			got := s.fn(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", s.name, got, tt.want)
			}
		}
	}
}
