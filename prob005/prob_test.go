package prob005

import (
	"testing"
)

func TestCar(t *testing.T) {
	if got, want := car(cons(3, 4)), 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCdr(t *testing.T) {
	got, want := cdr(cons(3, 4)), 4
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
