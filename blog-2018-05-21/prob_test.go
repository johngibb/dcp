package prob

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"10101010", "01010101"},
		{"00001011", "00000111"},
	}
	for _, tt := range tests {
		in, out := fromBinary(tt.in), fromBinary(tt.out)
		got := SwapEvenOddBits(in)
		if out != got {
			t.Errorf("got %s, want %s",
				toBinary(got),
				toBinary(out),
			)
		}
	}
}

func fromBinary(s string) byte {
	if len(s) > 8 {
		panic("too long")
	}
	var result byte
	p := byte(1)
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '1':
			result += p
		case '0':
		default:
			panic("unknown character: " + string(s[i]))
		}
		p *= 2
	}
	return result
}

func toBinary(b byte) string {
	var s string
	for i := 0; i < 8; i++ {
		if b&1 == 1 {
			s = "1" + s
		} else {
			s = "0" + s
		}
		b >>= 1
	}
	return s
}
