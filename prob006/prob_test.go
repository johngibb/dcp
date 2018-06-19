package prob006

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	head := &Node{Val: 1}
	for i := 2; i < 10; i++ {
		head.Add(i)
	}

	fmt.Println(head)

	assert := func(i, want int) {
		if got := head.Get(i); got != want {
			t.Errorf("head.Get(%d): got %d, want %d", i, got, want)
		}
	}
	assert(0, 1)
	assert(1, 2)
}
