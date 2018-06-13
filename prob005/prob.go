package prob005

// cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first
// and last element of that pair. For example, car(cons(3, 4)) returns 3, and
// cdr(cons(3, 4)) returns 4.
//
// Given this implementation of cons:
//
// def cons(a, b):
//     return lambda f : f(a, b)
// Implement car and cdr.
//
// This is pretty useless in go.

type pair func(fn func(int, int) int) int

func cons(a, b int) pair {
	return func(fn func(int, int) int) int {
		return fn(a, b)
	}
}

func car(p pair) int {
	return p(func(a, b int) int {
		return a
	})
}

func cdr(p pair) int {
	return p(func(a, b int) int {
		return b
	})
}
