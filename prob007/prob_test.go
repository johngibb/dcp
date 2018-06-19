package prob

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{"111", 3},
		{"126", 3},
		{"1", 1},
		{"11", 2},
		{"26", 2},
		{"27", 1},
		{"723", 2},
		{"2723", 2},
		{"27233", 2},
		{"27223", 3},
		{"2220", 2},
		{"", 0},
	}
	strategies := []struct {
		name string
		fn   func(string) int
	}{
		{"NumDecodes", NumDecodes},
		{"NumDecodes2", NumDecodes2},
	}
	for _, s := range strategies {
		t.Run(s.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.s, func(t *testing.T) {
					got := NumDecodes(tt.s)
					if got != tt.n {
						t.Errorf("got %d, want %d", got, tt.n)
					}
				})
			}
		})
	}
}
