package prob

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"aabcdef", "aa"},
		{"abccccd", "cccc"},
		{"aabcdcb", "bcdcb"},
		{"bananas", "anana"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := LongestPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s          string
		palindrome bool
	}{
		{"aba", true},
		{"a", true},
		{"ab", false},
		{"abba", true},
	}
	for _, tt := range tests {
		got := IsPalindrome(tt.s)
		if got != tt.palindrome {
			t.Errorf("%s: got %v, want %v", tt.s, got, tt.palindrome)
		}
	}
}
